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
	dhttp "github.com/Filimonov-ua-d/home_finance_new/finances/delivery/http"
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

	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	dbname := viper.GetString("postgres.dbname")
	sslmode := viper.GetString("postgres.sslmode")
	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.port")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		user, password, dbname, sslmode, host, port)

	if db, err = sqlx.Connect("postgres", dsn); err != nil {
		log.Panic("Error DB connect")
	}

	financeRepo := pg.NewFinancesRepository(db)

	return &App{
		financeUC: fnusecase.NewFinanceUseCase(financeRepo),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()

	dhttp.RegisterHTTPEndpoints(router, a.financeUC)

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
