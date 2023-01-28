package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Filimonov-ua-d/home_finance_new/finances"
	pg "github.com/Filimonov-ua-d/home_finance_new/finances/repository/postgres"
	fnusecase "github.com/Filimonov-ua-d/home_finance_new/finances/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server
	financeUC  finances.UseCase
}

func NewApp(db *sqlx.DB) *App {

	var err error

	dsn := viper.GetString("postgres.user")
	if db, err = sqlx.Connect("postgres", dsn); err != nil {
		fmt.Println(err)
	}

	financeRepo := pg.NewFinancesRepository(db)

	return &App{
		financeUC: fnusecase.NewFinanceUseCase(financeRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
