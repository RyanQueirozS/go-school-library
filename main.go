package main

import (
	"fmt"
	"go-school-library/models"
)

func main() {
	acc := models.Account{}
	acc.SetAccountId(1)
	acc.SetAccountName("name")
	fmt.Println(acc.GetID())
	fmt.Println(acc.GetName())

	// bk := methods.Book{iD: 1, name: "bkName", description: "bkDes"}
}
