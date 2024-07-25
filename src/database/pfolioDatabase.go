package database

import (
	//"crypto/sha256"

	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pichik/pfolio/src/data"
)

var createPfolioTableQuery = `
CREATE TABLE IF NOT EXISTS user (
	USERNAME TEXT,
	XTB_USD JSON,
	XTB_EUR JSON
);
`

var createUserQuery = `
  INSERT INTO user (
	USERNAME,
	XTB_USD,
	XTB_EUR
  ) VALUES (?,?,?);
`

func createUserDatabase() {
	_, err := userDb.Exec(createPfolioTableQuery)
	if err != nil {
		fmt.Printf("Error creating table in database: %s", err)
	}
}

func CreateUser(username string) {
	// Check if the user already exists
	var count int
	err := userDb.QueryRow("SELECT COUNT(*) FROM user WHERE USERNAME = ?", username).Scan(&count)
	if err != nil {
		fmt.Printf("Error checking if user exists: %s", err)
		return
	}

	if count > 0 {
		fmt.Println("User already exists")
		return
	}

	_, err = userDb.Exec(createUserQuery, username, "{}", "{}")

	if err != nil {
		fmt.Printf("Error adding data into database: %s", err)
	}
}

func AddPfolio(username string, xtb_JSON []byte, currency string) {

	if currency == "USD" {
		currency = "XTB_USD"
	} else if currency == "EUR" {
		currency = "XTB_EUR"
	} else {
		return
	}

	// If the record exists, update the existing JSON data
	sqlQuery := fmt.Sprintf("UPDATE user SET %s = ? WHERE USERNAME = ?", currency)
	_, err := userDb.Exec(sqlQuery, xtb_JSON, username)

	if err != nil {
		fmt.Println("Error updating/inserting record:", err)
		return
	}
}

func ReadPfolio(username string) ([]data.XTB_CSV, []data.XTB_CSV) {
	// Query to select XTB_USD and XTB_EUR columns from the user table
	query := `SELECT XTB_USD, XTB_EUR FROM user	WHERE USERNAME = ?;`
	var xtbUSD []data.XTB_CSV
	var xtbEUR []data.XTB_CSV

	// Execute the query
	row := userDb.QueryRow(query, username)
	var xtbUSDJSON, xtbEURJSON sql.NullString
	err := row.Scan(&xtbUSDJSON, &xtbEURJSON)
	if err != nil {
		fmt.Printf("error reading XTB data: %v\n", err)
	}

	// Check if XTB_USDJSON is not NULL before unmarshaling
	if xtbUSDJSON.Valid {
		err = json.Unmarshal([]byte(xtbUSDJSON.String), &xtbUSD)
		if err != nil {
			fmt.Printf("error unmarshalling XTB_USD JSON: %v\n", err)
		}
	}

	// Check if XTB_EURJSON is not NULL before unmarshaling
	if xtbEURJSON.Valid {
		err = json.Unmarshal([]byte(xtbEURJSON.String), &xtbEUR)
		if err != nil {
			fmt.Printf("error unmarshalling XTB_EUR JSON: %v\n", err)
		}
	}

	return xtbUSD, xtbEUR
}
