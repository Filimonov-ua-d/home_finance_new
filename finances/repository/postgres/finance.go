package postgres

import (
	"context"
	"time"

	"github.com/Filimonov-ua-d/home_finance_new/models"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type FinancesRepository struct {
	DB     *sqlx.DB
	logger *zerolog.Logger
}

func NewFinancesRepository(db *sqlx.DB, logger *zerolog.Logger) *FinancesRepository {
	return &FinancesRepository{
		DB:     db,
		logger: logger,
	}
}

func (fr *FinancesRepository) InsertProfit(ctx context.Context, p *models.Profit) error {

	var profit = toDBProfit(p)

	ctx = context.Background()
	tx, err := fr.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(ctx, "insert into profit (amount,source,date) VALUES (:amount,:source,:date)", &profit)
	if err != nil {

		tx.Rollback()

		fr.logger.Error().
			Err(err).
			Str("Func:", "InsertProfit(into profit)")

		return err
	}

	_, err = tx.NamedExecContext(ctx, "insert into money (amount,date) VALUES (:amount,:date)", &profit)
	if err != nil {

		tx.Rollback()

		fr.logger.Error().
			Err(err).
			Str("Func:", "InsertProfit(into money)")

		return err
	}

	err = tx.Commit()
	if err != nil {

		fr.logger.Error().
			Err(err).
			Str("Func:", "InsertProfit(tx.Commit)")

		return err
	}

	return nil
}

func (fr *FinancesRepository) InsertDeposit(ctx context.Context, d *models.Deposit) (err error) {

	var deposit = toDBDeposit(d)

	insertSQL := "insert into deposit (name,description,amount,returndate,date) VALUES (:name,:description,:amount,:returndate,:date)"

	_, err = fr.DB.NamedExec(insertSQL, deposit)
	if err != nil {

		fr.logger.Error().
			Err(err).
			Str("Func:", "InsertDeposit")

		return
	}

	return

}

func (fr *FinancesRepository) InsertSalary(ctx context.Context, s *models.Salary) (err error) {

	insertSQL := "insert into salary (amount,date) VALUES (:amount,:date)"

	var salary = toDBSalary(s)

	_, err = fr.DB.NamedExec(insertSQL, salary)
	if err != nil {

		fr.logger.Error().
			Err(err).
			Str("Func:", "InsertSalary")

		return
	}

	return

}

func (fr *FinancesRepository) InsertCredit(ctx context.Context, c *models.Credit) (err error) {

	insertSQL := "insert into credit (amount, date, name, description, returndate) VALUES (:amount,:date, :name, :description, :returndate)"

	var credit = toDBCredit(c)

	_, err = fr.DB.NamedExec(insertSQL, credit)
	if err != nil {

		fr.logger.Error().
			Err(err).
			Str("Func:", "InsertCredit")

		return
	}

	return

}

func (fr *FinancesRepository) InsertExpense(c context.Context, e *models.Expense) (err error) {

	insertSQL := "insert into expences (name,amount,date) VALUES (:name,:amount,:date)"

	var expense = toDBExpense(e)

	_, err = fr.DB.NamedExec(insertSQL, expense)
	if err != nil {

		fr.logger.Error().
			Err(err).
			Str("Func:", "InsertExpense")

		return
	}

	return

}

func (fr *FinancesRepository) InsertExpensesItem(ctx context.Context, ei *models.ExpensesItem) (err error) {

	insertSQL := "insert into expence_items (name) VALUES (:name)"

	var expensesItem = toDBExpensesItem(ei)

	_, err = fr.DB.NamedExec(insertSQL, expensesItem)
	if err != nil {

		fr.logger.Error().
			Err(err).
			Str("Func:", "InsertExpensesItem")

		return
	}

	return

}

func (fr *FinancesRepository) GetExpences(ctx context.Context) (exp []*models.ExpensesItem, err error) {

	selectSQL := "SELECT * FROM expence_items"

	var expences []*ExpensesItem

	if err = fr.DB.Select(&expences, selectSQL); err != nil {

		fr.logger.Error().
			Err(err).
			Str("Func:", "GetExpences")

		return
	}

	for _, v := range expences {
		exp = append(exp, toModelExpensesItem(v))

	}

	return

}

func (fr *FinancesRepository) GetMoney(ctx context.Context, date time.Time) (mitems []*models.MoneyItem, err error) {

	selectSQL := `select sum(m.amount) as amount, cast(max(date($1)) AS DATE) as "date" from money m where m.date <= $1`

	var moneyItems []*MoneyItem

	if err := fr.DB.Select(&moneyItems, selectSQL, date); err != nil {

		fr.logger.Error().
			Err(err).
			Str("Func:", "GetMoney")

		return nil, err
	}

	for _, m := range moneyItems {
		mitems = append(mitems, toModelMoneyItem(m))
	}

	return
}
