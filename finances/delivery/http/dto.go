package http

import (
	"time"
)

type GetMoneyResponse struct {
	Amount int    `json:"amount"`
	Date   string `json:"date"`
}

type Profit struct {
	Amount int    `json:"amount"`
	Source string `json:"source"`
	Date   string `json:"date"`
}

type Salary struct {
	Amount int    `json:"amount"`
	Date   string `json:"date"`
}

type Credit struct {
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ReturnDate  string `json:"returndate"`
}

type ExpensesItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Expense struct {
	Date        string `json:"date"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type Money struct {
	Amount int       `json:"amount"`
	Date   time.Time `json:"date"`
}

type Deposit struct {
	Date        string `json:"date"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	ReturnDate  string `json:"returndate"`
}
