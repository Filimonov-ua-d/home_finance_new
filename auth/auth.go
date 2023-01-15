package auth

import (
	"fmt"

	req "github.com/Filimonov-ua-d/home_finance_new/requests"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func auth() {
	var err error
	dsn := "user=postgres password=postgres dbname=home_finance sslmode=disable"
	if req.DB, err = sqlx.Connect("postgres", dsn); err != nil {
		fmt.Println(err)
	}
}
