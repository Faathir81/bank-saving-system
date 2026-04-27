package controllers

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"bank-saving-system/config"
	"bank-saving-system/models"
	"bank-saving-system/utils"

	"gorm.io/gorm"
)

type TransactionRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
	Date      string  `json:"date"` // Format: YYYY-MM-DD
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	req := new(TransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid input", "")
		return
	}
	if req.Amount <= 0 {
		utils.SendError(w, http.StatusBadRequest, "Amount must be greater than 0", "")
		return
	}

	txDate, _ := time.Parse("2006-01-02", req.Date)

	var account models.Account
	if err := config.DB.First(&account, "id = ?", req.AccountID).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Account not found", "")
		return
	}

	// USE GORM TRANSACTION FOR ACID COMPLIANCE
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Create Transaction Record
		transaction := models.Transaction{
			AccountID:       account.ID,
			Type:            "deposit",
			Amount:          req.Amount,
			TransactionDate: txDate,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// Update Account Balance
		account.Balance += req.Amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to process deposit transaction", err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]interface{}{"message": "Deposit successful", "new_balance": account.Balance})
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	req := new(TransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid input", "")
		return
	}
	if req.Amount <= 0 {
		utils.SendError(w, http.StatusBadRequest, "Amount must be greater than 0", "")
		return
	}

	withdrawDate, _ := time.Parse("2006-01-02", req.Date)

	var account models.Account
	if err := config.DB.Preload("DepositoType").First(&account, "id = ?", req.AccountID).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Account not found", "")
		return
	}

	if account.Balance < req.Amount {
		utils.SendError(w, http.StatusBadRequest, "Insufficient balance", "")
		return
	}

	// --- LOGIC PERHITUNGAN BUNGA (Sesuai Soal) ---
	// Kita asumsikan #months adalah selisih bulan dari tanggal pembuatan akun/deposit terakhir ke tanggal withdrawal
	// Untuk kemudahan tes, kita hitung selisih bulan antara created_at akun dan tanggal withdrawal yang diinput

	months := math.Floor(withdrawDate.Sub(account.CreatedAt).Hours() / 24 / 30)
	if months < 0 {
		months = 0
	}

	monthlyReturn := account.DepositoType.YearlyReturn / 12

	// Rumus soal: ending balance = starting balance * #months * monthly return
	// (Seperti diskusi kita, ini adalah nilai BUNGA yang didapat)
	interestEarned := account.Balance * months * monthlyReturn
	totalWithdrawal := req.Amount + interestEarned

	// Update Record
	// USE GORM TRANSACTION FOR ACID COMPLIANCE
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		transaction := models.Transaction{
			AccountID:       account.ID,
			Type:            "withdraw",
			Amount:          req.Amount,
			TransactionDate: withdrawDate,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		account.Balance -= req.Amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to process withdraw transaction", err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]interface{}{
		"message":          "Withdrawal successful",
		"starting_balance": account.Balance + req.Amount,
		"amount_withdrawn": req.Amount,
		"months_stayed":    months,
		"interest_earned":  interestEarned,
		"total_received":   totalWithdrawal,
		"current_balance":  account.Balance,
	})
}
