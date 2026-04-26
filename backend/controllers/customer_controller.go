package controllers

import (
	"encoding/json"
	"net/http"

	"bank-saving-system/config"
	"bank-saving-system/models"
	"bank-saving-system/utils"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer
	config.DB.Find(&customers)
	utils.SendJSON(w, http.StatusOK, customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var customer models.Customer
	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Customer not found", "")
		return
	}
	utils.SendJSON(w, http.StatusOK, customer)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	customer := new(models.Customer)
	if err := json.NewDecoder(r.Body).Decode(customer); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Review your input", err.Error())
		return
	}

	config.DB.Create(&customer)
	utils.SendJSON(w, http.StatusCreated, customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	customer := new(models.Customer)
	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Customer not found", "")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(customer); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Review your input", err.Error())
		return
	}

	config.DB.Save(&customer)
	utils.SendJSON(w, http.StatusOK, customer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var customer models.Customer
	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Customer not found", "")
		return
	}

	config.DB.Unscoped().Delete(&customer)
	utils.SendJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Customer deleted successfully"})
}
