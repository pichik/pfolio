package misc

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pichik/pfolio/src/data"
)

func CsvToJSON(csvData string) ([]byte, error) {
	// Create a CSV reader from the CSV data
	reader := csv.NewReader(strings.NewReader(csvData))
	reader.Comma = ';'

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return []byte{}, err
	}

	// Check if there are no records or only the header
	if len(records) <= 1 {
		return []byte("[]"), nil
	}

	// Create a slice to store the converted JSON data
	var jsonData []map[string]interface{}

	// Iterate through CSV records (skip the header) and convert to JSON
	for _, record := range records[1:] {
		jsonRecord := make(map[string]interface{})
		for i, value := range record {
			// Assuming the header row is present in the CSV,
			// use the header as the JSON key
			header := records[0][i]

			// Convert numeric values to float64, keep others as string
			if _, err := strconv.ParseFloat(value, 64); err == nil {
				jsonRecord[header], _ = strconv.ParseFloat(value, 64)
			} else {
				jsonRecord[header] = value
			}
		}
		jsonData = append(jsonData, jsonRecord)
	}

	// Convert JSON slice to JSON string
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return []byte{}, err
	}

	return jsonBytes, nil
}

func JsonToStockDataList(jsonData []byte) ([]data.XTBStockData, error) {
	// Create a slice to store the converted StockData
	var stockDataList []data.XTBStockData

	// Unmarshal the JSON data into a slice of maps
	var records []map[string]interface{}
	if err := json.Unmarshal(jsonData, &records); err != nil {
		return nil, err
	}

	// Iterate through JSON records and convert to StockData
	for _, record := range records {
		stockData := data.XTBStockData{
			ID:      fmt.Sprintf("%.0f", record["ID"].(float64)),
			Type:    record["Type"].(string),
			Time:    record["Time"].(string),
			Symbol:  strings.Split(record["Symbol"].(string), ".")[0],
			Comment: record["Comment"].(string),
			Amount:  record["Amount"].(float64),
		}

		stockDataList = append(stockDataList, stockData)
	}

	return stockDataList, nil
}
