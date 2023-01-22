package postgres

import (
	"time"

	"github.com/Filimonov-ua-d/home_finance_new/models"
)

type Profit struct {
	Amount int    `db:"amount"`
	Source string `db:"source"`
	Date   string `db:"date"`
}

type Salary struct {
	Amount int    `db:"amount"`
	Date   string `db:"date"`
}

type Credit struct {
	Amount      int    `db:"amount"`
	Date        string `db:"date"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ReturnDate  string `db:"returndate"`
}

type ExpensesItem struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Expense struct {
	Date        string `db:"date"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Amount      int    `db:"date"`
}

type MoneyItem struct {
	Amount int       `db:"amount"`
	Date   time.Time `db:"date"`
}

type Deposit struct {
	Date        string `db:"date"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Amount      int    `db:"amount"`
	ReturnDate  string `db:"returndate"`
}

func toDBProfit(p *models.Profit) *Profit {
	return &Profit{
		Amount: p.Amount,
		Source: p.Source,
		Date:   p.Date,
	}
}

func toDBDeposit(d *models.Deposit) *Deposit {
	return &Deposit{
		Date:        d.Date,
		Name:        d.Name,
		Description: d.Description,
		Amount:      d.Amount,
		ReturnDate:  d.ReturnDate,
	}
}

func toDBSalary(s *models.Salary) *Salary {
	return &Salary{
		Amount: s.Amount,
		Date:   s.Date,
	}
}

func toDBCredit(c *models.Credit) *Credit {
	return &Credit{
		Amount:      c.Amount,
		Date:        c.Date,
		Name:        c.Name,
		Description: c.Description,
		ReturnDate:  c.ReturnDate,
	}
}

func toDBExpensesItem(ei *models.ExpensesItem) *ExpensesItem {
	return &ExpensesItem{
		ID:   ei.ID,
		Name: ei.Name,
	}
}

func toDBExpense(e *models.Expense) *Expense {
	return &Expense{
		Date:        e.Date,
		Name:        e.Name,
		Description: e.Description,
		Amount:      e.Amount,
	}
}

func toModelExpensesItem(ei *ExpensesItem) *models.ExpensesItem {
	return &models.ExpensesItem{
		ID:   ei.ID,
		Name: ei.Name,
	}
}

func toModelMoneyItem(mi *MoneyItem) *models.MoneyItem {
	return &models.MoneyItem{
		Amount: mi.Amount,
		Date:   mi.Date,
	}
}
