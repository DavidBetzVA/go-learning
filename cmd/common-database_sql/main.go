package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb" // Import MSSQL driver
)

func main() {
	// Set up connection string
	connStr := "sqlserver://username:password@localhost:1433?database=mydb"
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer func() {
		_ = db.Close()
	}()

	// Query example
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("User:", id, name)
	}
}
