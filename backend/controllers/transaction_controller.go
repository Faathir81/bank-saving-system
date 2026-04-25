package controllers

import (
	"math"
	"time"

	"bank-saving-system/config"
	"bank-saving-system/models"

	"github.com/gofiber/fiber/v2"
)

type TransactionRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
	Date      string  `json:"date"` // Format: YYYY-MM-DD
}

func Deposit(c *fiber.Ctx) error {
	req := new(TransactionRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
	}

	txDate, _ := time.Parse("2006-01-02", req.Date)

	var account models.Account
	if err := config.DB.First(&account, "id = ?", req.AccountID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Account not found"})
	}

	// Create Transaction Record
	transaction := models.Transaction{
		AccountID:       account.ID,
		Type:            "deposit",
		Amount:          req.Amount,
		TransactionDate: txDate,
	}

	// Update Account Balance
	account.Balance += req.Amount

	config.DB.Create(&transaction)
	config.DB.Save(&account)

	return c.JSON(fiber.Map{"message": "Deposit successful", "new_balance": account.Balance})
}

func Withdraw(c *fiber.Ctx) error {
	req := new(TransactionRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input"})
	}

	withdrawDate, _ := time.Parse("2006-01-02", req.Date)

	var account models.Account
	if err := config.DB.Preload("DepositoType").First(&account, "id = ?", req.AccountID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Account not found"})
	}

	if account.Balance < req.Amount {
		return c.Status(400).JSON(fiber.Map{"message": "Insufficient balance"})
	}

	// --- LOGIC PERHITUNGAN BUNGA (Sesuai Soal) ---
	// Kita asumsikan #months adalah selisih bulan dari tanggal pembuatan akun/deposit terakhir ke tanggal withdrawal
	// Untuk kemudahan tes, kita hitung selisih bulan antara created_at akun dan tanggal withdrawal yang diinput
	
	months := math.Floor(withdrawDate.Sub(account.CreatedAt).Hours() / 24 / 30)
	if months < 0 { months = 0 }

	monthlyReturn := account.DepositoType.YearlyReturn / 12
	
	// Rumus soal: ending balance = starting balance * #months * monthly return
	// (Seperti diskusi kita, ini adalah nilai BUNGA yang didapat)
	interestEarned := account.Balance * months * monthlyReturn
	totalWithdrawal := req.Amount + interestEarned

	// Update Record
	transaction := models.Transaction{
		AccountID:       account.ID,
		Type:            "withdraw",
		Amount:          req.Amount,
		TransactionDate: withdrawDate,
	}

	account.Balance -= req.Amount

	config.DB.Create(&transaction)
	config.DB.Save(&account)

	return c.JSON(fiber.Map{
		"message":          "Withdrawal successful",
		"starting_balance": account.Balance + req.Amount,
		"amount_withdrawn": req.Amount,
		"months_stayed":    months,
		"interest_earned":  interestEarned,
		"total_received":   totalWithdrawal,
		"current_balance":  account.Balance,
	})
}
