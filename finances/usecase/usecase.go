package usecase

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

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

	err := f.financesRepo.InsertCredit(ctx, credit)

	log.Error().
		Err(err).
		Str("package:", "usecase").
		Str("Func:", "InsertCredit")

	return err
}

func (f *FinancesUseCase) InsertDeposit(ctx context.Context, d *models.Deposit) error {
	depo := &models.Deposit{
		Date:        d.Date,
		Name:        d.Name,
		Description: d.Description,
		ReturnDate:  d.ReturnDate,
		Amount:      d.Amount,
	}

	err := f.financesRepo.InsertDeposit(ctx, depo)

	log.Error().
		Err(err).
		Str("package:", "usecase").
		Str("Func:", "InsertDeposit")

	return err
}

func (f *FinancesUseCase) InsertExpense(ctx context.Context, e *models.Expense) error {
	exp := &models.Expense{
		Date:        e.Date,
		Name:        e.Name,
		Description: e.Description,
		Amount:      e.Amount,
	}

	err := f.financesRepo.InsertExpense(ctx, exp)

	log.Error().
		Err(err).
		Str("package:", "usecase").
		Str("Func:", "InsertExpense")

	return err
}

func (f *FinancesUseCase) InsertExpensesItem(ctx context.Context, ei *models.ExpensesItem) error {
	mei := &models.ExpensesItem{
		ID:   ei.ID,
		Name: ei.Name,
	}

	err := f.financesRepo.InsertExpensesItem(ctx, mei)

	log.Error().
		Err(err).
		Str("package:", "usecase").
		Str("Func:", "InsertExpensesItem")

	return err
}

func (f *FinancesUseCase) InsertProfit(ctx context.Context, p *models.Profit) error {
	prof := &models.Profit{
		Amount: p.Amount,
		Source: p.Source,
		Date:   p.Date,
	}

	err := f.financesRepo.InsertProfit(ctx, prof)

	log.Error().
		Err(err).
		Str("package:", "usecase").
		Str("Func:", "InsertProfit")

	return err

}

func (f *FinancesUseCase) InsertSalary(ctx context.Context, s *models.Salary) error {
	sal := &models.Salary{
		Amount: s.Amount,
		Date:   s.Date,
	}

	err := f.financesRepo.InsertSalary(ctx, sal)

	log.Error().
		Err(err).
		Str("package:", "usecase").
		Str("Func:", "InsertSalary")

	return err
}

func (f *FinancesUseCase) GetExpences(ctx context.Context) ([]*models.ExpensesItem, error) {
	c, err := f.financesRepo.GetExpences(ctx)

	log.Error().
		Err(err).
		Str("package:", "usecase").
		Str("Func:", "GetExpences")

	return c, err
}

func (f *FinancesUseCase) GetMoneyOnDate(ctx context.Context, date time.Time) ([]*models.MoneyItem, error) {
	c, err := f.financesRepo.GetMoney(ctx, date)

	log.Error().
		Err(err).
		Str("package:", "usecase").
		Str("Func:", "GetMoneyOnDate")

	return c, err
}
