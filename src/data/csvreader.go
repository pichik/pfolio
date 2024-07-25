package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
)

// Transaction represents a single transaction entry
type XTB_CSV struct {
	ID      int
	Type    string
	Time    string
	Symbol  string
	Comment string
	Amount  float64
}

func ProcessCSVFile(content multipart.File) ([]XTB_CSV, error) {
	reader := csv.NewReader(content)
	reader.Comma = ';'

	// Read and discard the header line
	_, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV header: %v", err)
	}

	var transactions []XTB_CSV
	lineNumber := 1 // Track the line number for better error reporting

	for {
		record, err := reader.Read()
		if err == io.EOF {
			// End of file reached, break out of the loop
			break
		}

		if err != nil || len(record) != 6 {
			continue
		}

		// Process the valid record
		id, _ := strconv.Atoi(record[0])
		amount, _ := strconv.ParseFloat(strings.Replace(record[5], ",", ".", 1), 64)

		transaction := XTB_CSV{
			ID:      id,
			Type:    record[1],
			Time:    record[2],
			Symbol:  record[3],
			Comment: record[4],
			Amount:  amount,
		}
		transactions = append(transactions, transaction)
		lineNumber++
	}

	return transactions, nil
}
