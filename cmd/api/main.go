package main

import (
	"fmt"

	"github.com/Filimonov-ua-d/home_finance_new/config"
	req "github.com/Filimonov-ua-d/home_finance_new/requests"
	"github.com/Filimonov-ua-d/home_finance_new/server"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	var db *sqlx.DB

	if err := config.Init(); err != nil {
		return
	}

	app := server.NewApp(db)

	if err := app.Run(viper.GetString("port")); err != nil {
		return
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
