package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:username@tcp(host:3306)/dbs")
	if err != nil {
		panic(err)
	}
	println(db)
	defer db.Close()

}
