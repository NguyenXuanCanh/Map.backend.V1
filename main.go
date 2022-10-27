package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/NguyenXuanCanh/go-starter/hello"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	Email    string
	FullName string
}

func main() {
	hello.SayHello()
	db, err := sql.Open("mysql", "root:xuancanh@tcp(127.0.0.1:3306)/fahasa")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT Email, FullName FROM ACCOUNT")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var account Account
		err := res.Scan(&account.Email, &account.FullName)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", account)
	}

}
