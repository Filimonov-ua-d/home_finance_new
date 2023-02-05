package http

import (
	"reflect"
	"testing"

	"github.com/Filimonov-ua-d/home_finance_new/finances"
	"github.com/gin-gonic/gin"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		useCase finances.UseCase
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.useCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_InsertDeposit(t *testing.T) {
	type fields struct {
		useCase finances.UseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				useCase: tt.fields.useCase,
			}
			h.InsertDeposit(tt.args.c)
		})
	}
}

func TestHandler_InsertExpense(t *testing.T) {
	type fields struct {
		useCase finances.UseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				useCase: tt.fields.useCase,
			}
			h.InsertExpense(tt.args.c)
		})
	}
}

func TestHandler_InsertExpensesItem(t *testing.T) {
	type fields struct {
		useCase finances.UseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				useCase: tt.fields.useCase,
			}
			h.InsertExpensesItem(tt.args.c)
		})
	}
}

func TestHandler_InsertCredit(t *testing.T) {
	type fields struct {
		useCase finances.UseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				useCase: tt.fields.useCase,
			}
			h.InsertCredit(tt.args.c)
		})
	}
}

func TestHandler_InsertSalary(t *testing.T) {
	type fields struct {
		useCase finances.UseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				useCase: tt.fields.useCase,
			}
			h.InsertSalary(tt.args.c)
		})
	}
}

func TestHandler_InsertProfit(t *testing.T) {
	type fields struct {
		useCase finances.UseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				useCase: tt.fields.useCase,
			}
			h.InsertProfit(tt.args.c)
		})
	}
}

func TestHandler_GetExpences(t *testing.T) {
	type fields struct {
		useCase finances.UseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				useCase: tt.fields.useCase,
			}
			h.GetExpences(tt.args.c)
		})
	}
}

func TestHandler_GetMoney(t *testing.T) {
	type fields struct {
		useCase finances.UseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				useCase: tt.fields.useCase,
			}
			h.GetMoney(tt.args.c)
		})
	}
}
