package datacenter

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func ImportXTB(w http.ResponseWriter, r *http.Request) {
	// Read the CSV data from the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	csvData := string(body)

	// Convert CSV to JSON
	jsonData, err := csvToJSON(csvData)
	if err != nil {
		http.Error(w, "Error converting CSV to JSON", http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	// response := map[string]interface{}{
	// 	"status":  "success",
	// 	"message": "Data received successfully",
	// }

	// Convert the response map to JSON
	// responseJSON, err := json.Marshal(jsonData)
	// if err != nil {
	// 	http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	// 	return
	// }

	// Set the Content-Type header to indicate JSON
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func csvToJSON(csvData string) ([]byte, error) {
	// Create a CSV reader from the CSV data
	reader := csv.NewReader(strings.NewReader(csvData))

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return []byte{}, err
	}

	// Create a slice to store the converted JSON data
	var jsonData []map[string]string

	// Iterate through CSV records and convert to JSON
	for _, record := range records {
		jsonRecord := make(map[string]string)
		for i, value := range record {
			// Assuming the header row is present in the CSV,
			// use the header as the JSON key
			header := records[0][i]
			jsonRecord[header] = value
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
