package main

import (
	"fmt"

	req "github.com/Filimonov-ua-d/home_finance_new/requests"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	dsn := "user=postgres password=postgres dbname=home_finance sslmode=disable"
	if req.DB, err = sqlx.Connect("postgres", dsn); err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	router.POST("/profit/insert", req.InsertProfit)
	router.POST("/salary/insert", req.InsertSalary)
	router.POST("/credit/insert", req.InsertCredit)
	router.POST("/expensesitem/insert", req.InsertExpensesItem)
	router.POST("/expense/insert", req.InsertExpense)
	router.POST("/deposit/insert", req.InsertDeposit)
	router.GET("/expenceItems", req.GetExpences)
	router.POST("/money", req.GetMoney)

	router.Run("localhost:8080")
	fmt.Println("s")
}
