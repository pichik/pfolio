package database

import (
	//"crypto/sha256"

	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pichik/pfolio/src/data"
)

var readQuery = `
  		SELECT * FROM stocks;
  	`

// var db *sql.DB

func ReadDatabase() {
	// Query the data from the table

	rows, err := stockDb.Query(readQuery)
	if err != nil {
		fmt.Printf("Error querying database: %s", err)
	}
	defer rows.Close()

	for rows.Next() {

		var hourlyPricesJSON string
		var dailyPricesJSON string
		stock := data.Stock{}
		err := rows.Scan(
			&stock.Ticker,
			&stock.LastPrice,
			&stock.LastPriceStr,
			&hourlyPricesJSON,
			&dailyPricesJSON,
		)

		//If LastPrice is null it means data were not filled yet
		if err != nil {
			data.AllStocks[stock.Ticker] = stock
			return
		}
		// Unmarshal the JSON string into a map
		// Unmarshal the JSON string into a map
		var hourlyPrices map[int]float64
		if err := json.Unmarshal([]byte(hourlyPricesJSON), &hourlyPrices); err != nil {
			data.AllStocks[stock.Ticker] = stock
			return
		}
		stock.HourlyPrices = hourlyPrices

		// Unmarshal the JSON string into a map
		var dailyPrices map[int]float64
		if err := json.Unmarshal([]byte(dailyPricesJSON), &dailyPrices); err != nil {
			data.AllStocks[stock.Ticker] = stock
			return
		}
		stock.DailyPrices = dailyPrices

		data.AllStocks[stock.Ticker] = stock
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error retrieving database row: %s", err)
	}
}
