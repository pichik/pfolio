package data

import (
	"strconv"
	"strings"

	"github.com/pichik/pfolio/src/misc"
)

func FillAllStock(stocks map[string]OwnedStock) []OwnedStock {
	var list []OwnedStock
	for t, v := range stocks {
		if val, ok := AllStocks[strings.Split(t, ".")[0]]; ok {
			v.Stock = val
		}
		list = append(list, v)
	}
	return list
}

func GetPortfolio(xtbcsv []XTB_CSV) map[string]OwnedStock {
	// lines := strings.Split(input, "\n")
	result := make(map[string]OwnedStock)

	//Get previous line for dividend tax

	for _, line := range xtbcsv {

		if strings.Contains(line.Type, "purchase") {
			getPurchases(line, &result)
		} else if strings.Contains(line.Type, "Dividend") {
			var taxl XTB_CSV
			for _, taxline := range xtbcsv {
				if taxline.ID == line.ID+1 {
					taxl = taxline
					break
				}
			}
			getDividends(line, taxl, &result)
		}
	}

	return result
}

func getPurchases(line XTB_CSV, result *map[string]OwnedStock) {
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

	purchaseData := PurchaseData{
		Timestamp: misc.GetTimeStamp(line.Time),
		Quantity:  quantity,
		Price:     price,
	}

	ownedStock := (*result)[line.Symbol]

	ownedStock.Ticker = line.Symbol
	ownedStock.BuyPrice = ((ownedStock.BuyAmount * ownedStock.BuyPrice) + (quantity * price)) / (ownedStock.BuyAmount + quantity)
	ownedStock.BuyAmount += quantity

	if ownedStock.Purchases == nil {
		ownedStock.Purchases = []PurchaseData{}
	}

	ownedStock.Purchases = append(ownedStock.Purchases, purchaseData)
	(*result)[line.Symbol] = ownedStock

}

func getDividends(line XTB_CSV, taxline XTB_CSV, result *map[string]OwnedStock) {

	tax := 0.0
	if taxline.Type == "Withholding tax" {
		tax = taxline.Amount
	}

	dividendData := DividendData{
		Timestamp:   misc.GetTimeStamp(line.Time),
		Payout:      line.Amount,
		Tax:         tax,
		TaxedPayout: line.Amount + tax,
	}

	ownedStock := (*result)[line.Symbol]
	ownedStock.Ticker = line.Symbol
	ownedStock.Dividend += (line.Amount + tax)

	// Initialize Dividends slice if it's nil
	if ownedStock.Dividends == nil {
		ownedStock.Dividends = []DividendData{}
	}

	ownedStock.Dividends = append(ownedStock.Dividends, dividendData)
	(*result)[line.Symbol] = ownedStock
}
