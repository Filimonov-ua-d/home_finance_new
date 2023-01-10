package main

import (
	_ req "github.com/Filimonov-ua-d/home_finance_new/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	dsn := "user=postgres password=postgres dbname=home_finance sslmode=disable"
	if DB, err = sqlx.Connect("postgres", dsn); err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	router.POST("/profit/insert", InsertProfit)
	router.POST("/salary/insert", InsertSalary)
	router.POST("/credit/insert", InsertCredit)
	router.POST("/expensesitem/insert", InsertExpensesItem)
	router.POST("/expense/insert", InsertExpense)
	router.POST("/deposit/insert", InsertDeposit)
	router.GET("/expenceItems", GetExpences)
	router.POST("/money", GetMoney)

	router.Run("localhost:8080")
	fmt.Println("s")
}
