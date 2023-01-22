package requests

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	unmarshal "github.com/Filimonov-ua-d/home_finance_new/unmarshal"
)

var (
	ExpenceItems []ExpensesItem
	Expence      []Expense
	Deposits     []Deposit
	Credits      []Credit
	Salaries     []Salary
	DB           *sqlx.DB
)

type GetMoneyResponse struct {
	Amount int    `json:"amount"`
	Date   string `json:"date"`
}

type Profit struct {
	Amount int    `json:"amount" db:"amount"`
	Source string `json:"source" db:"source"`
	Date   string `json:"date" db:"date"`
}

type Salary struct {
	Amount int    `json:"amount" db:"amount"`
	Date   string `json:"date" db:"date"`
}

type Credit struct {
	Amount      int    `json:"amount" db:"amount"`
	Date        string `json:"date" db:"date"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	ReturnDate  string `json:"returndate" db:"returndate"`
}

type ExpensesItem struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Expense struct {
	Date        string `json:"date" db:"date"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount" db:"amount"`
}

type Money struct {
	Amount int       `json:"amount" db:"amount"`
	Date   time.Time `json:"date" db:"date" `
}

type Deposit struct {
	Date        string `json:"date" db:"date"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Amount      int    `json:"amount" db:"amount"`
	ReturnDate  string `json:"returndate" db:"returndate"`
}

func InsertDeposit(c *gin.Context) {

	insertSQL := "insert into deposit (name,description,amount,returndate,date) VALUES (:name,:description,:amount,:returndate,:date)"

	fmt.Println("Inserting deposit")

	var deposit Deposit
	if err := c.ShouldBindJSON(&deposit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Deposits = append(Deposits, deposit)

	res, err := DB.NamedExec(insertSQL, deposit)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(deposit)
}

func InsertExpense(c *gin.Context) {

	insertSQL := "insert into expences (name,amount,date) VALUES (:name,:amount,:date)"

	fmt.Println("Inserting expense")

	var expense Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Expence = append(Expence, expense)

	res, err := DB.NamedExec(insertSQL, expense)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(expense)
}

func InsertExpensesItem(c *gin.Context) {
	insertSQL := "insert into expence_items (name) VALUES (:name)"

	fmt.Println("Inserting Item of expensions")

	var expensesItem ExpensesItem
	if err := c.ShouldBindJSON(&expensesItem); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	ExpenceItems = append(ExpenceItems, expensesItem)

	res, err := DB.NamedExec(insertSQL, expensesItem)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(ExpenceItems)
}

func InsertCredit(c *gin.Context) {

	insertSQL := "insert into credit (amount, date, name, description, returndate) VALUES (:amount,:date, :name, :description, :returndate)"

	fmt.Println("Inserting Credit")

	var credit Credit
	if err := c.ShouldBindJSON(&credit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Credits = append(Credits, credit)

	res, err := DB.NamedExec(insertSQL, credit)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(credit)
}

func InsertSalary(c *gin.Context) {

	insertSQL := "insert into salary (amount,date) VALUES (:amount,:date)"

	fmt.Println("Inserting Salary")

	var salary Salary
	if err := c.ShouldBindJSON(&salary); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Salaries = append(Salaries, salary)

	res, err := DB.NamedExec(insertSQL, salary)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(res.LastInsertId())
	fmt.Println(salary)
}

func InsertProfit(c *gin.Context) {

	var profit Profit

	if err := c.ShouldBindJSON(&profit); err != nil {
		fmt.Println("Error bind JSON on &profit")
		c.String(http.StatusBadRequest, err.Error())
	}

	ctx := context.Background()
	tx, err := DB.BeginTxx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	_, err = tx.NamedExecContext(ctx, "insert into profit (amount,source,date) VALUES (:amount,:source,:date)", &profit)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	_, err = tx.NamedExecContext(ctx, "insert into money (amount,date) VALUES (:amount,:date)", &profit)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(tx)
}

func GetMoney(c *gin.Context) {

	var req unmarshal.GetMoneyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())

		return
	}

	selectSQL := `select sum(m.amount) as amount, cast(max(date($1)) AS DATE) as "date" from money m where m.date <= $1`

	fmt.Println("Get Money")

	var money []Money

	if err := DB.Select(&money, selectSQL, req.Date); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var mresp []GetMoneyResponse
	for _, m := range money {
		mresp = append(mresp, GetMoneyResponse{
			Amount: m.Amount,
			Date:   m.Date.Format("02.01.2006"),
		},
		)
	}

	c.JSON(http.StatusOK, mresp)
}

func GetExpences(c *gin.Context) {

	selectSQL := "SELECT * FROM expence_items"

	fmt.Println("Get expence items")

	if err := DB.Select(&ExpenceItems, selectSQL); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, ExpenceItems)

}
