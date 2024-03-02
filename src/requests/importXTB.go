package request

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/pichik/pfolio/src/data"
	"github.com/pichik/pfolio/src/misc"
)

func ImportXTB(w http.ResponseWriter, r *http.Request) {

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max size
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing multipart form: %v", err), http.StatusInternalServerError)
		return
	}

	// Retrieve the form values
	currency := r.FormValue("xtb-currency")

	file, _, err := r.FormFile("xtb-file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading part: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the file content
	body, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file content: %v", err), http.StatusInternalServerError)
		return
	}

	csvData := string(body)

	tck, pf := GetPortfolio(csvData)

	sendTickers(tck)

	cookie := &http.Cookie{
		Name:   "XTB-" + currency,
		Value:  pf,
		MaxAge: 365 * 24 * 60 * 60, // 1 year in seconds
	}
	// Set the cookie in the response
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func GetPortfolio(input string) (string, string) {
	lines := strings.Split(input, "\n")
	result := make(map[string]data.OwnedStock)

	for _, line := range lines {

		ownedStockData := data.OwnedStock{}

		if strings.Contains(line, "OPEN") {
			fields := strings.Split(line, ";")
			// time := fields[0]
			ticker := fields[3]
			priceStr := strings.Split(fields[4], " ")[4]

			quantityStr := strings.Split(strings.Split(fields[4], " ")[2], "/")[0]

			quantity, err := strconv.ParseFloat(quantityStr, 64)
			if err != nil {
				misc.ErrorLog.Printf("Error quantity parsing:  %s", err)
				continue
			}
			price, err := strconv.ParseFloat(priceStr, 64)
			if err != nil {
				misc.ErrorLog.Printf("Error amount parsing:  %s", err)
				continue
			}

			//calculate average price if there is higher quantity
			if val, ok := result[ticker]; ok {
				ownedStockData.BuyAmount = val.BuyAmount + quantity
				ownedStockData.BuyPrice = ((val.BuyAmount * val.BuyPrice) + (quantity * price)) / (val.BuyAmount + quantity)
				ownedStockData.Dividend = val.Dividend

			} else {
				ownedStockData.BuyAmount = quantity
				ownedStockData.BuyPrice = price
			}
			result[ticker] = ownedStockData

		} else if strings.Contains(line, "Dividend") {
			fields := strings.Split(line, ";")
			ticker := fields[3]

			dividendStr := fields[5]

			dividend, err := strconv.ParseFloat(strings.TrimSuffix(dividendStr, "\r"), 64)
			if err != nil {
				misc.ErrorLog.Printf("Error dividend parsing:  %s", err)
				continue
			}
			if val, ok := result[ticker]; ok {

				ownedStockData.BuyAmount = val.BuyAmount
				ownedStockData.BuyPrice = val.BuyPrice
				ownedStockData.Dividend = val.Dividend + dividend

			} else {
				ownedStockData.Dividend = dividend
			}
			result[ticker] = ownedStockData
		}
	}

	var tck string
	var ptf string

	for k, v := range result {
		tck = tck + k + " "
		ptf = ptf + fmt.Sprintf("%s:%f:%f:%f/", k, v.BuyAmount, v.BuyPrice, v.Dividend)
	}

	return tck, ptf
}

// func CheckClosedStocks(lines []string) {

// 	for _, line := range lines {
// 		if strings.Contains(line, "CLOSED") {
// 			fields := strings.Split(line, ";")
// 			// time := fields[0]
// 			ticker := fields[3]
// 			priceStr := strings.Split(fields[4], " ")[4]

// 			quantityStr := strings.Split(strings.Split(fields[4], " ")[2], "/")[0]

// 			quantity, err := strconv.ParseFloat(quantityStr, 64)
// 			if err != nil {
// 				misc.ErrorLog.Printf("Error quantity parsing:  %s", err)
// 				continue
// 			}
// 			price, err := strconv.ParseFloat(priceStr, 64)
// 			if err != nil {
// 				misc.ErrorLog.Printf("Error amount parsing:  %s", err)
// 				continue
// 			}
// 		}
// 	}

// }
