package finances

import (
	"context"

	mod "github.com/Filimonov-ua-d/home_finance_new/models"
)

type Repository interface {
	InsertDeposit(ctx context.Context, d *mod.Deposit) error
	InsertExpense(ctx context.Context, e *mod.Expense) error
	InsertExpensesItem(ctx context.Context, ei *mod.ExpensesItem) error
	InsertCredit(ctx context.Context, c *mod.Credit) error
	InsertSalary(ctx context.Context, s *mod.Salary) error
	InsertProfit(ctx context.Context, p *mod.Profit) error
	GetMoneyResponse(ctx context.Context) ([]*mod.GetMoneyResponse, error)
	GetExpences(ctx context.Context) ([]*mod.Expense, error)
}
