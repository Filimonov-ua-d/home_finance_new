package http

import (
	"net/http"

	"github.com/Filimonov-ua-d/home_finance_new/finances"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase finances.UseCase
}

func NewHandler(useCase finances.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) InsertDeposit(c *gin.Context) {

	var deposit Deposit

	if err := c.ShouldBindJSON(&deposit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.useCase.InsertDeposit(c.Request.Context(), toModelDeposit(&deposit)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) InsertExpense(c *gin.Context) {

	var expense Expense

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.useCase.InsertExpense(c.Request.Context(), toModelExpense(&expense)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

func (h *Handler) InsertExpensesItem(c *gin.Context) {

	var expensesItem ExpensesItem

	if err := c.ShouldBindJSON(&expensesItem); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.useCase.InsertExpensesItem(c.Request.Context(), toModelExpenseItem(&expensesItem)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

func (h *Handler) InsertCredit(c *gin.Context) {

	var credit Credit

	if err := c.ShouldBindJSON(&credit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.useCase.InsertCredit(c.Request.Context(), toModelCredit(&credit)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)

}

func (h *Handler) InsertSalary(c *gin.Context) {

	var salary Salary

	if err := c.ShouldBindJSON(&salary); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.useCase.InsertSalary(c.Request.Context(), toModelSalary(&salary)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)

}

func (h *Handler) InsertProfit(c *gin.Context) {

	var profit Profit

	if err := c.ShouldBindJSON(&profit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.useCase.InsertProfit(c.Request.Context(), toModelProfit(&profit)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)

}

type getExpItemsResponse struct {
	Items []*ExpensesItem
}

func (h *Handler) GetExpences(c *gin.Context) {

	exp, err := h.useCase.GetExpences(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getExpItemsResponse{
		Items: toExpenceItems(exp),
	})
}

func (h *Handler) GetMoney(c *gin.Context) {

	var req GetMoneyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())

		return
	}

	sum, err := h.useCase.GetMoneyOnDate(c.Request.Context(), req.Date)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, toMoney(sum))
}
