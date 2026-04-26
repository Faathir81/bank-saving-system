package controllers

import (
	"bank-saving-system/config"
	"bank-saving-system/models"

	"github.com/gofiber/fiber/v2"
)

func GetDepositoTypes(c *fiber.Ctx) error {
	var types []models.DepositoType
	config.DB.Find(&types)
	return c.JSON(types)
}

func CreateDepositoType(c *fiber.Ctx) error {
	deposito := new(models.DepositoType)
	if err := c.BodyParser(deposito); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	config.DB.Create(&deposito)
	return c.Status(201).JSON(deposito)
}

// SeedDepositoTypes helps populate initial data
func SeedDepositoTypes(c *fiber.Ctx) error {
	types := []models.DepositoType{
		{Name: "Bronze", YearlyReturn: 0.03},
		{Name: "Silver", YearlyReturn: 0.05},
		{Name: "Gold", YearlyReturn: 0.07},
	}

	for i := range types {
		config.DB.Where(models.DepositoType{Name: types[i].Name}).FirstOrCreate(&types[i])
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Seeds planted!"})
}

// CleanupDuplicateDepositoTypes removes duplicate deposito types, keeping the oldest per name
func CleanupDuplicateDepositoTypes(c *fiber.Ctx) error {
	result := config.DB.Exec(`
		DELETE FROM deposito_types
		WHERE id NOT IN (
			SELECT DISTINCT ON (name) id
			FROM deposito_types
			ORDER BY name, created_at ASC
		)
	`)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Duplicates removed!", "rows_affected": result.RowsAffected})
}

func UpdateDepositoType(c *fiber.Ctx) error {
	id := c.Params("id")
	deposito := new(models.DepositoType)
	if err := config.DB.First(&deposito, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Deposito Type not found"})
	}

	if err := c.BodyParser(deposito); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	config.DB.Save(&deposito)
	return c.JSON(deposito)
}

func DeleteDepositoType(c *fiber.Ctx) error {
	id := c.Params("id")
	var deposito models.DepositoType
	if err := config.DB.First(&deposito, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Deposito Type not found"})
	}

	// Unscoped deletion as we don't have soft delete set up properly to cascade, but note that this might fail if accounts are using it.
	if err := config.DB.Unscoped().Delete(&deposito).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Cannot delete Deposito Type: it is likely in use by accounts."})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deposito Type deleted successfully"})
}

