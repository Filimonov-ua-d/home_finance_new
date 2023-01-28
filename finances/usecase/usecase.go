package usecase

import (
	"context"
	"time"

	"github.com/Filimonov-ua-d/home_finance_new/finances"
	"github.com/Filimonov-ua-d/home_finance_new/models"
)

type FinancesUseCase struct {
	financesRepo finances.Repository
}

func NewFinanceUseCase(financesRepo finances.Repository) *FinancesUseCase {
	return &FinancesUseCase{
		financesRepo: financesRepo,
	}
}

func (f *FinancesUseCase) InsertCredit(ctx context.Context, c *models.Credit) error {
	credit := &models.Credit{
		Amount:      c.Amount,
		Date:        c.Date,
		Name:        c.Name,
		Description: c.Description,
		ReturnDate:  c.ReturnDate,
	}

	return f.financesRepo.InsertCredit(ctx, credit)
}

func (f *FinancesUseCase) InsertDeposit(ctx context.Context, d *models.Deposit) error {
	depo := &models.Deposit{
		Date:        d.Date,
		Name:        d.Name,
		Description: d.Description,
		ReturnDate:  d.ReturnDate,
		Amount:      d.Amount,
	}

	return f.financesRepo.InsertDeposit(ctx, depo)
}

func (f *FinancesUseCase) InsertExpense(ctx context.Context, e *models.Expense) error {
	exp := &models.Expense{
		Date:        e.Date,
		Name:        e.Name,
		Description: e.Description,
		Amount:      e.Amount,
	}

	return f.financesRepo.InsertExpense(ctx, exp)
}

func (f *FinancesUseCase) InsertExpensesItem(ctx context.Context, ei *models.ExpensesItem) error {
	mei := &models.ExpensesItem{
		ID:   ei.ID,
		Name: ei.Name,
	}

	return f.financesRepo.InsertExpensesItem(ctx, mei)
}

func (f *FinancesUseCase) InsertProfit(ctx context.Context, p *models.Profit) error {
	prof := &models.Profit{
		Amount: p.Amount,
		Source: p.Source,
		Date:   p.Date,
	}

	return f.financesRepo.InsertProfit(ctx, prof)
}

func (f *FinancesUseCase) InsertSalary(ctx context.Context, s *models.Salary) error {
	sal := &models.Salary{
		Amount: s.Amount,
		Date:   s.Date,
	}

	return f.financesRepo.InsertSalary(ctx, sal)
}

func (f *FinancesUseCase) GetExpences(ctx context.Context) ([]*models.ExpensesItem, error) {
	return f.financesRepo.GetExpences(ctx)
}

func (f *FinancesUseCase) GetMoneyOnDate(ctx context.Context, date time.Time) ([]*models.MoneyItem, error) {
	return f.financesRepo.GetMoney(ctx, date)
}
