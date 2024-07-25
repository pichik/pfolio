package request

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

var UsdToEur float64
var EurToUsd float64

func UpdateEchangeRates() {
	new := getExchangeRates("USD", "EUR")
	if new != 0 {
		UsdToEur = new
	}
	new = getExchangeRates("EUR", "USD")
	if new != 0 {
		EurToUsd = new
	}

}

func getExchangeRates(from string, to string) float64 {
	url := fmt.Sprintf("https://www.google.com/finance/quote/%s-%s?ucbcb=1", from, to)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	exchange := getExchange(string(body))

	return exchange
}

func getExchange(body string) float64 {
	// Define start and end strings
	startStr := `<div class="YMlKec fxKbKc">`
	endStr := `</div>`

	pattern := startStr + `([^<]*)` + endStr
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(body)
	if matches == nil {
		return 0
	}

	priceFloat, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0
	}
	return priceFloat
}
