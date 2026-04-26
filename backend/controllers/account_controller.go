package controllers

import (
	"bank-saving-system/config"
	"bank-saving-system/models"

	"github.com/gofiber/fiber/v2"
)

func GetAccounts(c *fiber.Ctx) error {
	var accounts []models.Account
	config.DB.Preload("Customer").Preload("DepositoType").Find(&accounts)
	return c.JSON(accounts)
}

func CreateAccount(c *fiber.Ctx) error {
	account := new(models.Account)
	if err := c.BodyParser(account); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	// Validate if customer and deposito type exist
	var customer models.Customer
	if err := config.DB.First(&customer, "id = ?", account.CustomerID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found"})
	}

	var deposito models.DepositoType
	if err := config.DB.First(&deposito, "id = ?", account.DepositoTypeID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Deposito Type not found"})
	}

	config.DB.Create(&account)
	
	// Load relationships for response
	config.DB.Preload("Customer").Preload("DepositoType").First(&account)
	
	return c.Status(201).JSON(account)
}

func UpdateAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	var account models.Account
	if err := config.DB.First(&account, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}

	// We only allow updating the DepositoType for an account
	type UpdatePayload struct {
		DepositoTypeID string `json:"deposito_type_id"`
	}
	var payload UpdatePayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}

	var deposito models.DepositoType
	if err := config.DB.First(&deposito, "id = ?", payload.DepositoTypeID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Deposito Type not found"})
	}

	account.DepositoTypeID = payload.DepositoTypeID
	config.DB.Save(&account)
	
	config.DB.Preload("Customer").Preload("DepositoType").First(&account)
	return c.JSON(account)
}

func DeleteAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	var account models.Account
	if err := config.DB.First(&account, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Account not found"})
	}

	// Delete related transactions first to avoid FK constraint errors
	config.DB.Unscoped().Where("account_id = ?", id).Delete(&models.Transaction{})

	// Now delete the account itself
	if err := config.DB.Unscoped().Delete(&account).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Account deleted successfully"})
}
