package localcache

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertDeposit(c *gin.Context) {

	insertSQL := "insert into deposit (name,description,amount,returndate,date) VALUES (:name,:description,:amount,:returndate,:date)"

	fmt.Println("Inserting deposit")

	var deposit Deposit
	if err := c.ShouldBindJSON(&deposit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Deposits = append(Deposits, deposit)

	res, err := DB.NamedExec(insertSQL, deposit)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(deposit)
}
