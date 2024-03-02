package data

type XTBStockData struct {
	ID      string  `json:"ID"`
	Type    string  `json:"Type"`
	Time    string  `json:"Time"`
	Symbol  string  `json:"Symbol"`
	Comment string  `json:"Comment"`
	Amount  float64 `json:"Amount"`
}

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
}
