package http

import (
	"os"

	"github.com/Filimonov-ua-d/home_finance_new/finances"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc finances.UseCase) {

	loggerHandler := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("Layer:", "handler").
		Str("Service:", "Home_finances").
		Logger()

	h := NewHandler(uc, &loggerHandler)

	router.POST("/profit/insert", h.InsertProfit)
	router.POST("/salary/insert", h.InsertSalary)
	router.POST("/credit/insert", h.InsertCredit)
	router.POST("/expensesitem/insert", h.InsertExpensesItem)
	router.POST("/expense/insert", h.InsertExpense)
	router.POST("/deposit/insert", h.InsertDeposit)
	router.GET("/expenceItems", h.GetExpences)
	router.POST("/money", h.GetMoney)
}
