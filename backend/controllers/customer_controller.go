package controllers

import (
	"bank-saving-system/config"
	"bank-saving-system/models"

	"github.com/gofiber/fiber/v2"
)

func GetCustomers(c *fiber.Ctx) error {
	var customers []models.Customer
	config.DB.Find(&customers)
	return c.JSON(customers)
}

func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer
	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found", "data": nil})
	}
	return c.JSON(customer)
}

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)
	if err := c.BodyParser(customer); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	config.DB.Create(&customer)
	return c.Status(201).JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	customer := new(models.Customer)
	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found", "data": nil})
	}

	if err := c.BodyParser(customer); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	config.DB.Save(&customer)
	return c.JSON(customer)
}

func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer
	if err := config.DB.First(&customer, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customer not found", "data": nil})
	}

	config.DB.Unscoped().Delete(&customer)
	return c.JSON(fiber.Map{"status": "success", "message": "Customer deleted successfully"})
}
