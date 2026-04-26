package controllers

import (
	"encoding/json"
	"net/http"

	"bank-saving-system/config"
	"bank-saving-system/models"
	"bank-saving-system/utils"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	var accounts []models.Account
	config.DB.Preload("Customer").Preload("DepositoType").Find(&accounts)
	utils.SendJSON(w, http.StatusOK, accounts)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	account := new(models.Account)
	if err := json.NewDecoder(r.Body).Decode(account); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Review your input", err.Error())
		return
	}

	// Validate if customer and deposito type exist
	var customer models.Customer
	if err := config.DB.First(&customer, "id = ?", account.CustomerID).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Customer not found", "")
		return
	}

	var deposito models.DepositoType
	if err := config.DB.First(&deposito, "id = ?", account.DepositoTypeID).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Deposito Type not found", "")
		return
	}

	config.DB.Create(&account)

	// Load relationships for response
	config.DB.Preload("Customer").Preload("DepositoType").First(&account)

	utils.SendJSON(w, http.StatusCreated, account)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var account models.Account
	if err := config.DB.First(&account, "id = ?", id).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Account not found", "")
		return
	}

	// We only allow updating the DepositoType for an account
	type UpdatePayload struct {
		DepositoTypeID string `json:"deposito_type_id"`
	}
	var payload UpdatePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Review your input", "")
		return
	}

	var deposito models.DepositoType
	if err := config.DB.First(&deposito, "id = ?", payload.DepositoTypeID).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Deposito Type not found", "")
		return
	}

	account.DepositoTypeID = payload.DepositoTypeID
	config.DB.Save(&account)

	config.DB.Preload("Customer").Preload("DepositoType").First(&account)
	utils.SendJSON(w, http.StatusOK, account)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var account models.Account
	if err := config.DB.First(&account, "id = ?", id).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Account not found", "")
		return
	}

	// Delete related transactions first to avoid FK constraint errors
	config.DB.Unscoped().Where("account_id = ?", id).Delete(&models.Transaction{})

	// Now delete the account itself
	if err := config.DB.Unscoped().Delete(&account).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Account deleted successfully"})
}
