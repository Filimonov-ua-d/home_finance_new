package models

import (
	"time"
)

type Profit struct {
	Amount int
	Source string
	Date   string
}

type Salary struct {
	Amount int
	Date   string
}

type Credit struct {
	Amount      int
	Date        string
	Name        string
	Description string
	ReturnDate  string
}

type ExpensesItem struct {
	ID   int
	Name string
}

type Expense struct {
	Date        string
	Name        string
	Description string
	Amount      int
}

type MoneyItem struct {
	Amount int
	Date   time.Time
}

type Deposit struct {
	Date        string
	Name        string
	Description string
	Amount      int
	ReturnDate  string
}
