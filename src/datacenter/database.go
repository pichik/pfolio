package data

import (
	//"crypto/sha256"

	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pichik/pfolio/src/misc"
)

var readQuery = `
  		SELECT * FROM triggers
  		WHERE Ticker = ?;
  	`

var db *sql.DB

var dataPath string = misc.DataDir + "/" + "data"

func ReadDatabase() {
	// Query the data from the table

	rows, err := db.Query(readQuery)
	if err != nil {
		fmt.Printf("Error querying database: %s", err)
	}
	defer rows.Close()

	for rows.Next() {

		var hourlyPricesJSON string
		stock := Stock{}
		err := rows.Scan(
			&stock.Ticker,
			&stock.LastPrice,
			&hourlyPricesJSON,
		)

		//If LastPrice is null it means data were not filled yet
		if err != nil {
			AllStocks[stock.Ticker] = stock
			return
		}
		// Unmarshal the JSON string into a map
		var hourlyPrices map[int]float64
		if err := json.Unmarshal([]byte(hourlyPricesJSON), &hourlyPrices); err != nil {
			fmt.Printf("Error unmarshaling hourlyPrices JSON: %s", err)
			continue
		}
		stock.HourlyPrices = hourlyPrices
		fmt.Println("JSON: ", hourlyPrices)

		AllStocks[stock.Ticker] = stock

	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error retrieving database row: %s", err)
	}
}

func Opendb(database string) *sql.DB {

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		misc.ErrorLog.Printf("Error opening database: %s", err)
	}
	return db
}
