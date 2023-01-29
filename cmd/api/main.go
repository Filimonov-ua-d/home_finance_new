package main

import (
	"log"

	"github.com/Filimonov-ua-d/home_finance_new/config"
	"github.com/Filimonov-ua-d/home_finance_new/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	var db *sqlx.DB

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp(db)

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}

}
