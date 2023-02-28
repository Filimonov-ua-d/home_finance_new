package usecase

import (
	"context"
	"errors"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"

	mockRepo "github.com/Filimonov-ua-d/home_finance_new/finances/repository/mock"
	"github.com/Filimonov-ua-d/home_finance_new/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

func TestFinancesUseCase_InsertCredit(t *testing.T) {

	err := errors.New("test")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testRepo := mockRepo.NewMockRepository(mockController)
	testRepo.EXPECT().InsertCredit(gomock.Any(), gomock.Any()).AnyTimes()
	testRepo.EXPECT().InsertCredit(gomock.Any(), gomock.Any()).AnyTimes().Return(err)

	type args struct {
		ctx context.Context
		c   *models.Credit
	}
	tests := []struct {
		name    string
		f       *FinancesUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			f: &FinancesUseCase{
				financesRepo: testRepo,
				logger:       &logger,
			},
			args: args{
				c: &models.Credit{
					Amount:      2222,
					Date:        "12.10.2005",
					Name:        "gotest",
					Description: "gotest",
					ReturnDate:  "12.11.2005",
				},
				ctx: ctx,
			},
			wantErr: false,
		},

		{
			name: "test2",
			f: &FinancesUseCase{
				financesRepo: testRepo,
				logger:       &logger,
			},
			args: args{
				ctx: ctx,
				c: &models.Credit{
					Amount:      2222,
					Date:        "12.10.2005",
					Name:        "gotest2",
					Description: "gotest2",
					ReturnDate:  "12.11.2005",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.InsertCredit(tt.args.ctx, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FinancesUseCase.InsertCredit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFinancesUseCase_InsertDeposit(t *testing.T) {

	err := errors.New("test")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testRepo := mockRepo.NewMockRepository(mockController)
	testRepo.EXPECT().InsertDeposit(gomock.Any(), gomock.Any()).AnyTimes().Return(err)

	type args struct {
		ctx context.Context
		d   *models.Deposit
	}
	tests := []struct {
		name    string
		f       *FinancesUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			f: &FinancesUseCase{
				financesRepo: testRepo,
				logger:       &logger,
			},
			args: args{
				ctx: ctx,
				d: &models.Deposit{
					Date:        "12.10.2005",
					Name:        "gotest2",
					Description: "gotest2",
					Amount:      2222,
					ReturnDate:  "12.11.2005",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.InsertDeposit(tt.args.ctx, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("FinancesUseCase.InsertDeposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFinancesUseCase_InsertExpense(t *testing.T) {

	err := errors.New("test")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testRepo := mockRepo.NewMockRepository(mockController)
	testRepo.EXPECT().InsertExpense(gomock.Any(), gomock.Any()).AnyTimes().Return(err)

	type args struct {
		ctx context.Context
		e   *models.Expense
	}
	tests := []struct {
		name    string
		f       *FinancesUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			f: &FinancesUseCase{
				financesRepo: testRepo,
				logger:       &logger,
			},
			args: args{
				ctx: ctx,
				e: &models.Expense{
					Date:        "12.10.2005",
					Name:        "gotest2",
					Description: "gotest2",
					Amount:      2222,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.InsertExpense(tt.args.ctx, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("FinancesUseCase.InsertExpense() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFinancesUseCase_InsertExpensesItem(t *testing.T) {

	err := errors.New("test")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testRepo := mockRepo.NewMockRepository(mockController)
	testRepo.EXPECT().InsertExpensesItem(gomock.Any(), gomock.Any()).AnyTimes().Return(err)

	type args struct {
		ctx context.Context
		ei  *models.ExpensesItem
	}
	tests := []struct {
		name    string
		f       *FinancesUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			f: &FinancesUseCase{
				financesRepo: testRepo,
				logger:       &logger,
			},
			args: args{
				ctx: ctx,
				ei: &models.ExpensesItem{
					ID:   1,
					Name: "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.InsertExpensesItem(tt.args.ctx, tt.args.ei); (err != nil) != tt.wantErr {
				t.Errorf("FinancesUseCase.InsertExpensesItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFinancesUseCase_InsertProfit(t *testing.T) {

	err := errors.New("test")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testRepo := mockRepo.NewMockRepository(mockController)
	testRepo.EXPECT().InsertProfit(gomock.Any(), gomock.Any()).AnyTimes().Return(err)

	type args struct {
		ctx context.Context
		p   *models.Profit
	}
	tests := []struct {
		name    string
		f       *FinancesUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			f: &FinancesUseCase{
				financesRepo: testRepo,
				logger:       &logger,
			},
			args: args{
				ctx: ctx,
				p: &models.Profit{
					Amount: 2222,
					Source: "test",
					Date:   "12.05.2005",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.InsertProfit(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("FinancesUseCase.InsertProfit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFinancesUseCase_InsertSalary(t *testing.T) {

	err := errors.New("test")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testRepo := mockRepo.NewMockRepository(mockController)
	testRepo.EXPECT().InsertSalary(gomock.Any(), gomock.Any()).AnyTimes().Return(err)

	type args struct {
		ctx context.Context
		s   *models.Salary
	}
	tests := []struct {
		name    string
		f       *FinancesUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			f: &FinancesUseCase{
				financesRepo: testRepo,
				logger:       &logger,
			},
			args: args{
				ctx: ctx,
				s: &models.Salary{
					Amount: 2222,
					Date:   "12.05.2005",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.InsertSalary(tt.args.ctx, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("FinancesUseCase.InsertSalary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFinancesUseCase_GetExpences(t *testing.T) {

	mod := []*models.ExpensesItem{
		{
			ID:   1,
			Name: "test",
		},
	}

	err := errors.New("test")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testRepo := mockRepo.NewMockRepository(mockController)
	testRepo.EXPECT().GetExpences(gomock.Any()).AnyTimes().Return(mod, err)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		f       *FinancesUseCase
		args    args
		want    []*models.ExpensesItem
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			f: &FinancesUseCase{
				financesRepo: testRepo,
				logger:       &logger,
			},
			args: args{
				ctx: ctx,
			},
			want:    mod,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.GetExpences(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FinancesUseCase.GetExpences() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FinancesUseCase.GetExpences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFinancesUseCase_GetMoneyOnDate(t *testing.T) {

	mod := []*models.MoneyItem{
		{
			Amount: 2222,
			Date:   time.Now().AddDate(2005, 05, 12),
		},
	}

	err := errors.New("test")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testRepo := mockRepo.NewMockRepository(mockController)
	testRepo.EXPECT().GetMoney(gomock.Any(), gomock.Any()).AnyTimes().Return(mod, err)

	type args struct {
		ctx  context.Context
		date time.Time
	}
	tests := []struct {
		name    string
		f       *FinancesUseCase
		args    args
		want    []*models.MoneyItem
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			f: &FinancesUseCase{
				financesRepo: testRepo,
				logger:       &logger,
			},
			args: args{
				ctx:  ctx,
				date: time.Now().AddDate(2005, 05, 12),
			},
			want:    mod,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.GetMoneyOnDate(tt.args.ctx, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("FinancesUseCase.GetMoneyOnDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FinancesUseCase.GetMoneyOnDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
