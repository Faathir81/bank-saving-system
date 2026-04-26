package routes

import (
	"net/http"

	"bank-saving-system/controllers"
)

func SetupRoutes(mux *http.ServeMux) {
	// Customer Routes
	mux.HandleFunc("GET /api/customers", controllers.GetCustomers)
	mux.HandleFunc("GET /api/customers/{id}", controllers.GetCustomer)
	mux.HandleFunc("POST /api/customers", controllers.CreateCustomer)
	mux.HandleFunc("PUT /api/customers/{id}", controllers.UpdateCustomer)
	mux.HandleFunc("DELETE /api/customers/{id}", controllers.DeleteCustomer)

	// Deposito Type Routes
	mux.HandleFunc("GET /api/deposito-types", controllers.GetDepositoTypes)
	mux.HandleFunc("POST /api/deposito-types", controllers.CreateDepositoType)
	mux.HandleFunc("PUT /api/deposito-types/{id}", controllers.UpdateDepositoType)
	mux.HandleFunc("DELETE /api/deposito-types/{id}", controllers.DeleteDepositoType)
	mux.HandleFunc("POST /api/deposito-types/seed", controllers.SeedDepositoTypes)
	mux.HandleFunc("DELETE /api/deposito-types/cleanup-duplicates", controllers.CleanupDuplicateDepositoTypes)

	// Account Routes
	mux.HandleFunc("GET /api/accounts", controllers.GetAccounts)
	mux.HandleFunc("POST /api/accounts", controllers.CreateAccount)
	mux.HandleFunc("PUT /api/accounts/{id}", controllers.UpdateAccount)
	mux.HandleFunc("DELETE /api/accounts/{id}", controllers.DeleteAccount)

	// Transaction Routes
	mux.HandleFunc("POST /api/transactions/deposit", controllers.Deposit)
	mux.HandleFunc("POST /api/transactions/withdraw", controllers.Withdraw)
}
