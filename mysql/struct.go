package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type TableDesc struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	res, _ := db.Query("show tables")
	var table string
	for res.Next() {
		res.Scan(&table)
		fmt.Println("table:", table)

		desc, _ := db.Query("desc " + table)
		for desc.Next() {
			var td TableDesc
			desc.Scan(&td.Field, &td.Type, &td.Null, &td.Key, &td.Default, &td.Extra)
			fmt.Println(td)
		}
	}
}
