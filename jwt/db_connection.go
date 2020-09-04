package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (dbc *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "basil"
	dbName := "testdb"
	dbc, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	fmt.Println("Connected")
	if err != nil {
		panic(err.Error())
	}
	return dbc
}
