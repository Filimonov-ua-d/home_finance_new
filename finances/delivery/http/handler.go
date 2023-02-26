package http

import (
	"net/http"

	"github.com/Filimonov-ua-d/home_finance_new/finances"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Handler struct {
	useCase finances.UseCase
	logger  *zerolog.Logger
}

func NewHandler(useCase finances.UseCase, logger *zerolog.Logger) *Handler {
	return &Handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *Handler) InsertDeposit(c *gin.Context) {

	var deposit Deposit

	if err := c.ShouldBindJSON(&deposit); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertDeposit")

		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.InsertDeposit(c.Request.Context(), toModelDeposit(&deposit)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertDeposit")

		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) InsertExpense(c *gin.Context) {

	var expense Expense

	if err := c.ShouldBindJSON(&expense); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertExpense")

		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.InsertExpense(c.Request.Context(), toModelExpense(&expense)); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertExpense")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) InsertExpensesItem(c *gin.Context) {

	var expensesItem ExpensesItem

	if err := c.ShouldBindJSON(&expensesItem); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertExpensesItem")

		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.InsertExpensesItem(c.Request.Context(), toModelExpenseItem(&expensesItem)); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertExpensesItem")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) InsertCredit(c *gin.Context) {

	var credit Credit

	if err := c.ShouldBindJSON(&credit); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertCredit")

		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.InsertCredit(c.Request.Context(), toModelCredit(&credit)); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertCredit")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}

func (h *Handler) InsertSalary(c *gin.Context) {

	var salary Salary

	if err := c.ShouldBindJSON(&salary); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertSalary")

		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.InsertSalary(c.Request.Context(), toModelSalary(&salary)); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertSalary")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}

func (h *Handler) InsertProfit(c *gin.Context) {

	var profit Profit

	if err := c.ShouldBindJSON(&profit); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertProfit")

		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.useCase.InsertProfit(c.Request.Context(), toModelProfit(&profit)); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "InsertProfit")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}

type getExpItemsResponse struct {
	Items []*ExpensesItem
}

func (h *Handler) GetExpences(c *gin.Context) {

	exp, err := h.useCase.GetExpences(c.Request.Context())
	if err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "GetExpences")

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

		h.logger.Error().
			Err(err).
			Str("Func:", "GetMoney")

		return
	}

	sum, err := h.useCase.GetMoneyOnDate(c.Request.Context(), req.Date)
	if err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "GetMoney")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, toMoney(sum))
}
