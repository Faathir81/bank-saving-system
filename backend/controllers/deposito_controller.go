package controllers

import (
	"encoding/json"
	"net/http"

	"bank-saving-system/config"
	"bank-saving-system/models"
	"bank-saving-system/utils"
)

func GetDepositoTypes(w http.ResponseWriter, r *http.Request) {
	var types []models.DepositoType
	config.DB.Find(&types)
	utils.SendJSON(w, http.StatusOK, types)
}

func CreateDepositoType(w http.ResponseWriter, r *http.Request) {
	deposito := new(models.DepositoType)
	if err := json.NewDecoder(r.Body).Decode(deposito); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Review your input", err.Error())
		return
	}

	config.DB.Create(&deposito)
	utils.SendJSON(w, http.StatusCreated, deposito)
}

// SeedDepositoTypes helps populate initial data
func SeedDepositoTypes(w http.ResponseWriter, r *http.Request) {
	types := []models.DepositoType{
		{Name: "Bronze", YearlyReturn: 0.03},
		{Name: "Silver", YearlyReturn: 0.05},
		{Name: "Gold", YearlyReturn: 0.07},
	}

	for i := range types {
		config.DB.Where(models.DepositoType{Name: types[i].Name}).FirstOrCreate(&types[i])
	}

	utils.SendJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Seeds planted!"})
}

// CleanupDuplicateDepositoTypes removes duplicate deposito types, keeping the oldest per name
func CleanupDuplicateDepositoTypes(w http.ResponseWriter, r *http.Request) {
	result := config.DB.Exec(`
		DELETE FROM deposito_types
		WHERE id NOT IN (
			SELECT DISTINCT ON (name) id
			FROM deposito_types
			ORDER BY name, created_at ASC
		)
	`)
	if result.Error != nil {
		utils.SendError(w, http.StatusInternalServerError, result.Error.Error(), "")
		return
	}
	utils.SendJSON(w, http.StatusOK, map[string]interface{}{"status": "success", "message": "Duplicates removed!", "rows_affected": result.RowsAffected})
}

func UpdateDepositoType(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	deposito := new(models.DepositoType)
	if err := config.DB.First(&deposito, "id = ?", id).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Deposito Type not found", "")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(deposito); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Review your input", err.Error())
		return
	}

	config.DB.Save(&deposito)
	utils.SendJSON(w, http.StatusOK, deposito)
}

func DeleteDepositoType(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var deposito models.DepositoType
	if err := config.DB.First(&deposito, "id = ?", id).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Deposito Type not found", "")
		return
	}

	// Unscoped deletion as we don't have soft delete set up properly to cascade, but note that this might fail if accounts are using it.
	if err := config.DB.Unscoped().Delete(&deposito).Error; err != nil {
		utils.SendError(w, http.StatusBadRequest, "Cannot delete Deposito Type: it is likely in use by accounts.", "")
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Deposito Type deleted successfully"})
}
