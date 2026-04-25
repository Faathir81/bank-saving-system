# System Specification - Bank Saving System

## 1. Mockup / User Journey
![Dashboard Mockup](./assets/mockup.png)

### User Flow:
1. **Admin/User** logs in to the dashboard.
2. **Customer Management**: Add or edit customer data.
3. **Account Opening**: A customer opens an account and chooses a Deposito Type (Bronze/3%, Silver/5%, Gold/7%).
4. **Transactions**: 
   - User adds balance (Deposit).
   - User withdraws balance. 
5. **Auto Calculation**: System calculates interest earned based on the number of months the money stayed in the account.

## 2. Database Design (ERD)

```mermaid
erDiagram
    CUSTOMER ||--o{ ACCOUNT : "owns"
    DEPOSITO_TYPE ||--o{ ACCOUNT : "defines"
    ACCOUNT ||--o{ TRANSACTION : "records"

    CUSTOMER {
        string id PK
        string name
        datetime created_at
    }

    DEPOSITO_TYPE {
        string id PK
        string name
        float yearly_return
        datetime created_at
    }

    ACCOUNT {
        string id PK
        string customer_id FK
        string deposito_type_id FK
        float balance
        datetime created_at
    }

    TRANSACTION {
        string id PK
        string account_id FK
        string type "deposit / withdraw"
        float amount
        datetime transaction_date
        datetime created_at
    }
```

## 3. UML Diagrams

### Use Case Diagram
```mermaid
graph TD
    subgraph "Bank Saving System"
        UC1[Manage Customers]
        UC2[Manage Deposito Types]
        UC3[Open/Close Accounts]
        UC4[Process Deposit]
        UC5[Process Withdrawal]
        UC6[Calculate Interest]
    end

    Admin((Admin/User)) --> UC1
    Admin --> UC2
    Admin --> UC3
    Admin --> UC4
    Admin --> UC5
    UC5 -.->|include| UC6
```

### Class Diagram (MVC Pattern)
```mermaid
classDiagram
    class CustomerController {
        +GetCustomers()
        +CreateCustomer()
    }
    class AccountController {
        +CreateAccount()
        +GetAccounts()
    }
    class TransactionController {
        +Deposit()
        +Withdraw()
    }
    class CustomerModel {
        +String ID
        +String Name
    }
    class AccountModel {
        +String ID
        +Float Balance
        +String CustomerID
    }
    
    CustomerController ..> CustomerModel : "Uses"
    AccountController ..> AccountModel : "Uses"
    TransactionController ..> AccountModel : "Updates"
```

## 4. API Screen Mapping
| Screen | Action | API Endpoint |
| :--- | :--- | :--- |
| **Dashboard** | Page Load | `GET /api/customers`, `GET /api/accounts` |
| **Customer List** | Add Customer | `POST /api/customers` |
| **Account Management** | Open Account | `POST /api/accounts` |
| **Transaction Modal** | Deposit | `POST /api/transactions/deposit` |
| **Transaction Modal** | Withdraw | `POST /api/transactions/withdraw` |

## 5. API Specifications

| Endpoint | Method | Description |
| :--- | :--- | :--- |
| `/api/customers` | GET | List all customers |
| `/api/customers` | POST | Create new customer |
| `/api/deposito-types` | GET | List deposito types |
| `/api/accounts` | POST | Create new account |
| `/api/transactions/deposit` | POST | Add balance to account |
| `/api/transactions/withdraw` | POST | Withdraw & calc interest |

## 5. Error Handling & Edge Cases

### Edge Cases:
- **Insufficient Balance**: System must reject withdrawal if amount > current balance.
- **Minimum Stay**: What happens if withdrawal occurs in < 1 month? (We should define if it gets 0 interest or pro-rated).
- **Duplicate IDs**: Ensure UUIDs or unique constraints for Customer & Account IDs.
- **Negative Input**: Reject any deposit/withdrawal with negative amount.

### Error Codes:
- `400 Bad Request`: Validation errors (missing fields, negative amount).
- `404 Not Found`: Customer or Account not found.
- `422 Unprocessable Entity`: Business logic errors (e.g., withdrawing more than balance).
- `500 Internal Server Error`: Database connection issues.
