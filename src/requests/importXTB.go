package request

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"

	"github.com/pichik/pfolio/src/data"
	"github.com/pichik/pfolio/src/misc"
)

func ImportXTB(w http.ResponseWriter, r *http.Request) {

	// Get the multipart reader from the request
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, "Error getting multipart reader", http.StatusInternalServerError)
		return
	}

	// Read the first part of the multipart request
	part, err := reader.NextPart()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading part: %v", err), http.StatusInternalServerError)
		return
	}

	// Check if the part is a file
	if part.FileName() == "" {
		return
	}
	// Read the file content
	body, err := io.ReadAll(part)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file content: %v", err), http.StatusInternalServerError)
		return
	}

	csvData := string(body)

	// Convert CSV to JSON
	jsonData, err := misc.CsvToJSON(csvData)
	if err != nil {

		http.Error(w, fmt.Sprintf("Error converting CSV to JSON [%s]", err), http.StatusInternalServerError)
		return
	}

	// Convert JSON to StockData
	stockDataList, err := misc.JsonToStockDataList(jsonData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting JSON to StockData [%s]", err), http.StatusInternalServerError)
		return
	}

	portfolio := calculateTickers(stockDataList)

	var tickers string
	for t, _ := range portfolio {
		tickers = tickers + t + " "
	}

	sendTickers(tickers)

	// Convert stockDataList to JSON
	jsonResponse, err := json.Marshal(data.AllStocks)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting StockDataList to JSON [%s]", err), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func calculateTickers(portfolioData []XTBStockData) map[string]float64 {
	var tickers = map[string]float64{}

	for _, data := range portfolioData {
		if data.Symbol != "" {
			tickers[data.Symbol] += math.Abs(data.Amount)
		}
	}
	return tickers
}
