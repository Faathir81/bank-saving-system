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

	for _, t := range types {
		config.DB.Where(models.DepositoType{Name: t.Name}).FirstOrCreate(&t)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Seeds planted!"})
}
