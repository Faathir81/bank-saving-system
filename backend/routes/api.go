package routes

import (
	"bank-saving-system/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Customer Routes
	customer := api.Group("/customers")
	customer.Get("/", controllers.GetCustomers)
	customer.Get("/:id", controllers.GetCustomer)
	customer.Post("/", controllers.CreateCustomer)
	customer.Put("/:id", controllers.UpdateCustomer)
	customer.Delete("/:id", controllers.DeleteCustomer)

	// Deposito Type Routes
	deposito := api.Group("/deposito-types")
	deposito.Get("/", controllers.GetDepositoTypes)
	deposito.Post("/", controllers.CreateDepositoType)
	deposito.Post("/seed", controllers.SeedDepositoTypes)

	// Account Routes
	account := api.Group("/accounts")
	account.Get("/", controllers.GetAccounts)
	account.Post("/", controllers.CreateAccount)

	// Transaction Routes
	transaction := api.Group("/transactions")
	transaction.Post("/deposit", controllers.Deposit)
	transaction.Post("/withdraw", controllers.Withdraw)
}
