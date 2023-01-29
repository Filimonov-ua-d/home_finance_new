package http

import (
	"github.com/Filimonov-ua-d/home_finance_new/finances"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc finances.UseCase) {
	h := NewHandler(uc)

	router.POST("/profit/insert", h.InsertProfit)
	router.POST("/salary/insert", h.InsertSalary)
	router.POST("/credit/insert", h.InsertCredit)
	router.POST("/expensesitem/insert", h.InsertExpensesItem)
	router.POST("/expense/insert", h.InsertExpense)
	router.POST("/deposit/insert", h.InsertDeposit)
	router.GET("/expenceItems", h.GetExpences)
	router.POST("/money", h.GetMoney)
}
