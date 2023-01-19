package http

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/Filimonov-ua-d/home_finance_new/models"
)

type GetMoneyResponse struct {
	Moneys []*MoneyItem `json:"moneys"`
}

type MoneyItem struct {
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

type Deposit struct {
	Date        string `json:"date"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	ReturnDate  string `json:"returndate"`
}

type GetMoneyRequest struct {
	Date time.Time `json:"date" db:"date" tformat:"02.01.2006"`
	//Sum  string `json:"sum" db:"sum"`
}

func (d *GetMoneyRequest) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	rawDate := v.(map[string]interface{})
	vv := rawDate["date"].(string)

	c := reflect.TypeOf(*d).Field(0).Tag
	g := c.Get("tformat")

	dd, err := time.ParseInLocation(g, vv, time.Local)
	fmt.Println("Unmarshal rsult: ", dd, err)
	d.Date = dd
	return err
}

func toModelDeposit(d *Deposit) *models.Deposit {
	return &models.Deposit{
		Date:        d.Date,
		Name:        d.Name,
		Description: d.Description,
		Amount:      d.Amount,
		ReturnDate:  d.ReturnDate,
	}
}

func toModelExpense(e *Expense) *models.Expense {
	return &models.Expense{
		Date:        e.Date,
		Name:        e.Name,
		Description: e.Description,
		Amount:      e.Amount,
	}
}

func toModelExpenseItem(ei *ExpensesItem) *models.ExpensesItem {
	return &models.ExpensesItem{
		ID:   ei.ID,
		Name: ei.Name,
	}
}

func toModelCredit(c *Credit) *models.Credit {
	return &models.Credit{
		Date:        c.Date,
		Name:        c.Name,
		Description: c.Description,
		Amount:      c.Amount,
		ReturnDate:  c.ReturnDate,
	}
}

func toModelSalary(s *Salary) *models.Salary {
	return &models.Salary{
		Amount: s.Amount,
		Date:   s.Date,
	}
}

func toModelProfit(p *Profit) *models.Profit {
	return &models.Profit{
		Amount: p.Amount,
		Source: p.Source,
		Date:   p.Date,
	}
}

func toExpenceItem(exp *models.ExpensesItem) *ExpensesItem {
	return &ExpensesItem{
		ID:   exp.ID,
		Name: exp.Name,
	}
}

func toExpenceItems(exps []*models.ExpensesItem) []*ExpensesItem {

	var expenseItems []*ExpensesItem

	for _, v := range exps {
		expenseItems = append(expenseItems, toExpenceItem(v))
	}

	return expenseItems
}

func toMoney(m []*models.MoneyItem) *GetMoneyResponse {

	var money []*MoneyItem

	for _, v := range m {
		money = append(money, toMoneyItem(v))

	}

	return &GetMoneyResponse{
		Moneys: money,
	}

}

func toMoneyItem(m *models.MoneyItem) *MoneyItem {
	return &MoneyItem{
		Amount: m.Amount,
		Date:   m.Date.Format("02.01.2006"),
	}
}
