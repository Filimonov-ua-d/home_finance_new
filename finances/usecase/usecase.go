package usecase

import (
	"context"

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

func (f *FinancesUseCase) InsertCredit(ctx context.Context, amount int, date, name, description, returndate string) error {
	credit := &models.Credit{
		Amount:      amount,
		Date:        date,
		Name:        name,
		Description: description,
		ReturnDate:  returndate,
	}

	return f.financesRepo.InsertCredit(ctx, credit)
}

func (f *FinancesUseCase) InsertDeposit(ctx context.Context, amount int, date, name, description, returndate string) error {
	depo := &models.Deposit{
		Date:        date,
		Name:        name,
		Description: description,
		ReturnDate:  returndate,
		Amount:      amount,
	}

	return f.financesRepo.InsertDeposit(ctx, depo)
}

func (f *FinancesUseCase) InsertExpense(ctx context.Context, amount int, date, name, description string) error {
	exp := &models.Expense{
		Date:        date,
		Name:        name,
		Description: description,
		Amount:      amount,
	}

	return f.financesRepo.InsertExpense(ctx, exp)
}

func (f *FinancesUseCase) InsertExpensesItem(ctx context.Context, id int, name string) error {
	ei := &models.ExpensesItem{
		ID:   id,
		Name: name,
	}

	return f.financesRepo.InsertExpensesItem(ctx, ei)
}

func (f *FinancesUseCase) InsertProfit(ctx context.Context, amount int, source, date string) error {
	prof := &models.Profit{
		Amount: amount,
		Source: source,
		Date:   date,
	}

	return f.financesRepo.InsertProfit(ctx, prof)
}

func (f *FinancesUseCase) InsertSalary(ctx context.Context, amount int, date string) error {
	sal := &models.Salary{
		Amount: amount,
		Date:   date,
	}

	return f.financesRepo.InsertSalary(ctx, sal)
}

func (f *FinancesUseCase) GetExpences(ctx context.Context) ([]*models.Expense, error) {
	return f.financesRepo.GetExpences(ctx)
}

func (f *FinancesUseCase) GetMoneyOnDate(ctx context.Context) ([]*models.MoneyItem, error) {
	return f.financesRepo.GetMoneyResponse(ctx)
}
