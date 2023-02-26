package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	mockUC "github.com/Filimonov-ua-d/home_finance_new/finances/usecase/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func MockJsonPost(c *gin.Context /* the test context */, content interface{}) {
	c.Request.Method = "POST" // or PUT
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestHandler_InsertDeposit(t *testing.T) {

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(ctx, map[string]interface{}{
		"date":        "12.10.2005",
		"name":        "gotest",
		"description": "gotest",
		"amount":      2222,
		"returndate":  "12.11.2005"})
	//assert.EqualValues(t, http.StatusOK, w.Code)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testUC := mockUC.NewMockUseCase(mockController)
	testUC.EXPECT().InsertDeposit(gomock.Any(), gomock.Any()).AnyTimes()

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			h: &Handler{
				useCase: testUC,
				logger:  &logger,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.InsertDeposit(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}

func TestHandler_InsertExpense(t *testing.T) {

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(ctx, map[string]interface{}{
		"date":        "12.10.2005",
		"name":        "gotest",
		"description": "gotest",
		"amount":      2222,
	})

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testUC := mockUC.NewMockUseCase(mockController)
	testUC.EXPECT().InsertExpense(gomock.Any(), gomock.Any()).AnyTimes()

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test2",
			h: &Handler{
				useCase: testUC,
				logger:  &logger,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.InsertExpense(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}

func TestHandler_InsertExpensesItem(t *testing.T) {

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(ctx, map[string]interface{}{
		"id":   1,
		"name": "gotest",
	})

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testUC := mockUC.NewMockUseCase(mockController)
	testUC.EXPECT().InsertExpensesItem(gomock.Any(), gomock.Any()).AnyTimes()

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test3",
			h: &Handler{
				useCase: testUC,
				logger:  &logger,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.InsertExpensesItem(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}

func TestHandler_InsertCredit(t *testing.T) {

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testUC := mockUC.NewMockUseCase(mockController)
	testUC.EXPECT().InsertCredit(gomock.Any(), gomock.Any()).AnyTimes()

	MockJsonPost(ctx, map[string]interface{}{
		"amount":      222,
		"date":        "12.10.2005",
		"name":        "gotest",
		"description": "gotest",
		"returndate":  "12.11.2005",
	})

	type args struct {
		c *gin.Context
	}

	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "gotest4",
			h: &Handler{
				useCase: testUC,
				logger:  &logger,
			},
			args: args{
				c: ctx,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.InsertCredit(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}

func TestHandler_InsertSalary(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testUC := mockUC.NewMockUseCase(mockController)
	testUC.EXPECT().InsertSalary(gomock.Any(), gomock.Any()).AnyTimes()

	MockJsonPost(ctx, map[string]interface{}{
		"amount": 2222,
		"date":   "12.10.2005",
	})

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test4",
			h: &Handler{
				useCase: testUC,
				logger:  &logger,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.InsertSalary(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}

func TestHandler_InsertProfit(t *testing.T) {

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testUC := mockUC.NewMockUseCase(mockController)
	testUC.EXPECT().InsertProfit(gomock.Any(), gomock.Any()).AnyTimes()

	MockJsonPost(ctx, map[string]interface{}{
		"amount": 2222,
		"source": "gotest",
		"date":   "12.10.2005",
	})

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test6",
			h: &Handler{
				useCase: testUC,
				logger:  &logger,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.InsertProfit(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}

func TestHandler_GetExpences(t *testing.T) {

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testUC := mockUC.NewMockUseCase(mockController)
	testUC.EXPECT().GetExpences(gomock.Any()).AnyTimes()

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test7",
			h: &Handler{
				useCase: testUC,
				logger:  &logger,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetExpences(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}

func TestHandler_GetMoney(t *testing.T) {

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	testUC := mockUC.NewMockUseCase(mockController)
	testUC.EXPECT().GetMoneyOnDate(gomock.Any(), gomock.Any()).AnyTimes()

	MockJsonPost(ctx, map[string]interface{}{
		"date": "12.10.2005",
	})

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test8",
			h: &Handler{
				useCase: testUC,
				logger:  &logger,
			},
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetMoney(tt.args.c)
			assert.EqualValues(t, http.StatusOK, w.Code)
		})
	}
}
