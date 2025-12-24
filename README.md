# Golang CRUD Todo API

This is a CLI application for expense tracker to manage your expenses. 

Project from: https://roadmap.sh/projects/expense-tracker-api


## Features

- Login
- Register
- Refresh Token
- Create Expense
- Update Expense
- Delete Expense
- Get All Expense
- Get All Expense With Time Filter
- Get All Expense With Category Filter

## Documentation

### Register
Endpoint
```
POST /login
```

Request
```
{
    "name": "john doe",
    "email": "johndoe@gmail.com",
    "password": "password"
}
```

Response
```
HTTP 201/Created
{
  "message": "register success"
}
```

### Login
Endpoint
```
POST /login
```

Request
```
{
  "name": "John Doe",
  "email": "john@doe.com",
  "password": "password"
}
```

Response
```
HTTP 200/OK
{
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjY1NDU1NTIsInR5cGUiOiJhY2Nlc3MiLCJ1c2VySWQiOjJ9.j2G0DT1sN4Avvwi8vKDqN__gw2m64OL-p48fss7431U",
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjcxNDk0NTIsInR5cGUiOiJyZWZyZXNoIiwidXNlcklkIjoyfQ.sAb0fAq9LGSscBtFcNQQ_P20Wk2ekQn7emdaF-wnc6s"
}
```

### Create Expense
Endpoint
```
POST /expenses
```

Request
```
{
  "title": "Buy bill",
  "description": "Buy credit card",
  "amount": 20.0,
  "category": "Payment"
}
```

Response
```
HTTP 201/Created
{
    "ID": 4,
    "Title": "Buy bill",
    "Description": "Buy credit card",
    "Amount": 20,
    "Category": "Payment",
    "UserID": 2,
    "CreatedAt": "2025-12-24T09:48:08.286+07:00",
    "UpdatedAt": "2025-12-24T09:48:08.286+07:00"
}
```

### Update Expense
Endpoint
```
PUT /expenses/1
```

Request
```
{
  "title": "Buy bill",
  "description": "Buy credit card, netflix",
  "amount": 20.0,
  "category": "Payment"
}
```

Response
```
HTTP 200/OK
{
    "ID": 3,
    "Title": "Buy groceries",
    "Description": "Buy credit card, netflix",
    "Amount": 20,
    "Category": "Payment",
    "UserID": 2,
    "CreatedAt": "2025-12-24T09:47:34.581+07:00",
    "UpdatedAt": "2025-12-24T09:55:07.799+07:00"
}
```

## Delete Expense
Endpoint
```
DELETE /expenses/1
```

Request
```
-
```

Response
```
HTTP 204/No Content
-
```

## Get All Expenses
Endpoint
```
GET /expenses
```

Request
```
-
```

Response
```
{
    "data": [
        {
            "ID": 4,
            "Title": "Buy bill",
            "Description": "Buy credit card",
            "Amount": 20,
            "Category": "Payment",
            "UserID": 2,
            "CreatedAt": "2025-12-24T09:48:08.286+07:00",
            "UpdatedAt": "2025-12-24T09:48:08.286+07:00"
        },
        {
            "ID": 3,
            "Title": "Buy groceries",
            "Description": "Buy credit card, netflix",
            "Amount": 20,
            "Category": "Payment",
            "UserID": 2,
            "CreatedAt": "2025-12-24T09:47:34.581+07:00",
            "UpdatedAt": "2025-12-24T09:55:07.799+07:00"
        }
    ],
    "total": 40
}
```

## Get All Expense With Time Filter
Endpoint
```
GET /expenses?filterType=lastWeek
GET /expenses?filterType=lastMonth
GET /expenses?filterType=lastThreeMonth
```

Request
```
-
```

Response
```
{
    "data": [
        {
            "ID": 4,
            "Title": "Buy bill",
            "Description": "Buy credit card",
            "Amount": 20,
            "Category": "Payment",
            "UserID": 2,
            "CreatedAt": "2025-12-24T09:48:08.286+07:00",
            "UpdatedAt": "2025-12-24T09:48:08.286+07:00"
        },
        {
            "ID": 3,
            "Title": "Buy groceries",
            "Description": "Buy credit card, netflix",
            "Amount": 20,
            "Category": "Payment",
            "UserID": 2,
            "CreatedAt": "2025-12-24T09:47:34.581+07:00",
            "UpdatedAt": "2025-12-24T09:55:07.799+07:00"
        }
    ],
    "total": 40
}
```

## Get All Expense With Category Filter
Endpoint
```
GET /expenses?category=Payment
```

Request
```
-
```

Response
```
{
    "data": [
        {
            "ID": 4,
            "Title": "Buy bill",
            "Description": "Buy credit card",
            "Amount": 20,
            "Category": "Payment",
            "UserID": 2,
            "CreatedAt": "2025-12-24T09:48:08.286+07:00",
            "UpdatedAt": "2025-12-24T09:48:08.286+07:00"
        },
        {
            "ID": 3,
            "Title": "Buy groceries",
            "Description": "Buy credit card, netflix",
            "Amount": 20,
            "Category": "Payment",
            "UserID": 2,
            "CreatedAt": "2025-12-24T09:47:34.581+07:00",
            "UpdatedAt": "2025-12-24T09:55:07.799+07:00"
        }
    ],
    "total": 40
}
```


## Clone the project

```bash
git clone https://github.com/mathiasyeremiaaryadi/project-expense-tracker-crud-api-golang.git
```

Go to the project directory

```bash
cd project-expense-tracker-crud-api-golang
```

Install dependencies

```bash
go build
```

Start the server

```bash
./expense-tracker-api
```
