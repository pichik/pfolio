package request

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/pichik/pfolio/src/data"
)

func StocksUpdate(w http.ResponseWriter, r *http.Request) {

	// Set the Content-Type header to indicate JSON
	w.Header().Set("Content-Type", "application/json")

	// Create a map to hold all items
	response := map[string]interface{}{
		"UsdToEur": UsdToEur,
		"EurToUsd": EurToUsd,
	}

	xtbEurCookie, _ := r.Cookie("XTB-EUR")

	if xtbEurCookie != nil {
		xtbeur := CookiesToPortfolio(xtbEurCookie.Value)
		response["xtb_eur"] = xtbeur
	}

	xtbUsdCookie, _ := r.Cookie("XTB-USD")

	if xtbUsdCookie != nil {
		xtbusd := CookiesToPortfolio(xtbUsdCookie.Value)
		response["xtb_usd"] = xtbusd
	}

	wListCookie, _ := r.Cookie("wList")

	if wListCookie != nil {
		wList := CookiesToWlist(wListCookie.Value)
		response["wList"] = wList
	}

	// Convert stockDataList to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CookiesToPortfolio(cookie string) []data.OwnedStock {
	trimmed := strings.TrimSuffix(cookie, "/")
	stockEntries := strings.Split(trimmed, "/")

	var ownedStocks []data.OwnedStock

	// Iterate over each stock entry
	for _, entry := range stockEntries {
		// Split each entry by colon to separate ticker and price
		parts := strings.Split(entry, ":")

		ticker := parts[0]
		amount, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil
		}

		price, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil
		}

		dividend, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			return nil
		}

		// Create OwnedStock struct and append to slice
		ownedStock := data.OwnedStock{
			Ticker:    ticker,
			BuyPrice:  price,
			BuyAmount: amount,
			Dividend:  dividend,
		}

		if val, ok := data.AllStocks[strings.Split(ticker, ".")[0]]; ok {
			ownedStock.Stock = val
		}
		ownedStocks = append(ownedStocks, ownedStock)

	}

	return ownedStocks
}

func CookiesToWlist(cookie string) []data.Stock {
	stockEntries := strings.Split(cookie, ",")

	var stocks []data.Stock

	// Iterate over each stock entry
	for _, s := range stockEntries {
		// Split each entry by colon to separate ticker and price

		if val, ok := data.AllStocks[s]; ok {
			stocks = append(stocks, val)
		}
	}

	return stocks
}
