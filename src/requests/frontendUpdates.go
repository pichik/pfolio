package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/pichik/pfolio/src/data"
	"github.com/pichik/pfolio/src/database"
	"github.com/pichik/pfolio/src/misc"
)

func StocksUpdate(w http.ResponseWriter, r *http.Request) {

	// Set the Content-Type header to indicate JSON
	w.Header().Set("Content-Type", "application/json")

	// Create a map to hold all items
	response := map[string]interface{}{
		"UsdToEur": UsdToEur,
		"EurToUsd": EurToUsd,
	}

	username, err := r.Cookie("username")
	if err != nil {
		fmt.Println("Cookie not found:", err)
		return
	}
	usd, eur := database.ReadPfolio(username.Value)

	response["xtb_usd"] = AddStockData(GetPortfolio(usd))
	response["xtb_eur"] = AddStockData(GetPortfolio(eur))

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

func AddStockData(stocks map[string]data.OwnedStock) []data.OwnedStock {
	var list []data.OwnedStock
	for t, v := range stocks {
		if val, ok := data.AllStocks[strings.Split(t, ".")[0]]; ok {
			v.Stock = val
		}
		list = append(list, v)
	}
	return list
}

func GetPortfolio(xtbcsv []data.XTB_CSV) map[string]data.OwnedStock {
	// lines := strings.Split(input, "\n")
	result := make(map[string]data.OwnedStock)

	//Get previous line for dividend tax

	for _, line := range xtbcsv {

		if strings.Contains(line.Type, "purchase") {
			GetPurchases(line, &result)
		} else if strings.Contains(line.Type, "Dividend") {
			var taxl data.XTB_CSV
			for _, taxline := range xtbcsv {
				if taxline.ID == line.ID+1 {
					taxl = taxline
					break
				}
			}
			GetDividends(line, taxl, &result)
		}
	}

	return result
}

func GetPurchases(line data.XTB_CSV, result *map[string]data.OwnedStock) {
	priceStr := strings.Split(line.Comment, " ")[4]
	quantityStr := strings.Split(strings.Split(line.Comment, " ")[2], "/")[0]

	quantity, err := strconv.ParseFloat(quantityStr, 64)
	if err != nil {
		misc.ErrorLog.Printf("Error quantity parsing:  %s", err)
		return
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		misc.ErrorLog.Printf("Error amount parsing:  %s", err)
		return
	}

	purchaseData := data.PurchaseData{
		Timestamp: getTimeStamp(line.Time),
		Quantity:  quantity,
		Price:     price,
	}

	ownedStock := (*result)[line.Symbol]

	ownedStock.Ticker = line.Symbol
	ownedStock.BuyPrice = ((ownedStock.BuyAmount * ownedStock.BuyPrice) + (quantity * price)) / (ownedStock.BuyAmount + quantity)
	ownedStock.BuyAmount += quantity

	if ownedStock.Purchases == nil {
		ownedStock.Purchases = []data.PurchaseData{}
	}

	ownedStock.Purchases = append(ownedStock.Purchases, purchaseData)
	(*result)[line.Symbol] = ownedStock

}

func GetDividends(line data.XTB_CSV, taxline data.XTB_CSV, result *map[string]data.OwnedStock) {

	tax := 0.0
	if taxline.Type == "Withholding tax" {
		tax = taxline.Amount
	}

	dividendData := data.DividendData{
		Timestamp:   getTimeStamp(line.Time),
		Payout:      line.Amount,
		Tax:         tax,
		TaxedPayout: line.Amount + tax,
	}

	ownedStock := (*result)[line.Symbol]
	ownedStock.Ticker = line.Symbol
	ownedStock.Dividend += (line.Amount + tax)

	// Initialize Dividends slice if it's nil
	if ownedStock.Dividends == nil {
		ownedStock.Dividends = []data.DividendData{}
	}

	ownedStock.Dividends = append(ownedStock.Dividends, dividendData)
	(*result)[line.Symbol] = ownedStock
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
