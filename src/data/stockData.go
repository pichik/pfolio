package data

type Stock struct {
	Ticker             string
	LastPrice          float64
	LastPriceStr       string
	LastPriceTimestamp int
	HourlyPrices       map[int]float64
	DailyPrices        map[int]float64
}

var AllStocks = make(map[string]Stock)

type OwnedStock struct {
	Ticker    string
	Stock     Stock
	BuyAmount float64
	BuyPrice  float64
	Dividend  float64
	Purchases []PurchaseData
	Dividends []DividendData
}

type PurchaseData struct {
	Timestamp int64
	Quantity  float64
	Price     float64
}
type DividendData struct {
	Timestamp   int64
	Payout      float64
	Tax         float64
	TaxedPayout float64
}
