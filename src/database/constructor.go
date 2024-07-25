package database

import (
	"database/sql"

	"github.com/pichik/pfolio/src/misc"
)

var userDb *sql.DB
var stockDb *sql.DB

func Opendb(directory string) {
	opendb(directory+"userdata.db", &userDb)
	createUserDatabase()
	opendb(directory+"stockdata.db", &stockDb)
}

func opendb(database string, db **sql.DB) {
	var err error
	*db, err = sql.Open("sqlite3", database)
	if err != nil {
		misc.ErrorLog.Printf("Error opening database: %s", err)
	}
}
