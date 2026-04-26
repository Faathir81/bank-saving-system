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
	deposito.Put("/:id", controllers.UpdateDepositoType)
	deposito.Delete("/:id", controllers.DeleteDepositoType)
	deposito.Post("/seed", controllers.SeedDepositoTypes)
	deposito.Delete("/cleanup-duplicates", controllers.CleanupDuplicateDepositoTypes)

	// Account Routes
	account := api.Group("/accounts")
	account.Get("/", controllers.GetAccounts)
	account.Post("/", controllers.CreateAccount)
	account.Put("/:id", controllers.UpdateAccount)
	account.Delete("/:id", controllers.DeleteAccount)

	// Transaction Routes
	transaction := api.Group("/transactions")
	transaction.Post("/deposit", controllers.Deposit)
	transaction.Post("/withdraw", controllers.Withdraw)
}
