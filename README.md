# Bank Saving System

A Full-Stack Technical Test project for Intern position. This system manages customers, accounts, and deposits with interest calculation features.

## Tech Stack

- **Backend**: Golang (Fiber Framework)
- **Frontend**: React.js + Tailwind CSS
- **Database**: PostgreSQL
- **Documentation**: Postman, Mermaid UML

## Project Structure

```text
├── backend/          # Golang source code
├── frontend/         # React.js source code
├── docs/             # Documentation, UML, and API Specs
└── README.md
```

## Features

- CRUD Customer
- CRUD Account
- CRUD Deposito Type (Bronze, Silver, Gold)
- Deposit & Withdrawal Transactions
- Automatic Interest Calculation on Withdrawal

## Getting Started

## 1. Setup Database

- Install PostgreSQL
- Create a new database named `bank_saving_db`
- *Note: Database migration is handled automatically when you start the backend server.*

## 2. Setup Backend

```bash
# Install dependencies
cd backend
go mod tidy

# Run server
go run main.go
```

## 3. Setup Frontend

```bash
# Install dependencies
cd frontend
npm install

# Run server
npm run dev
```

## 4. Seeding Initial Data

Once both servers are running, you need to seed the initial Deposito Packages (Bronze, Silver, Gold).
You can do this by sending a `POST` request to:
`http://localhost:8080/api/deposito-types/seed`

Or simply run this in your terminal (PowerShell):
```powershell
Invoke-RestMethod -Method POST -Uri "http://localhost:8080/api/deposito-types/seed"
```
