package request

import (
	"bytes"
	"fmt"
	"net/http"
)

var tickerDbUrl = "http://localhost:8081/addticker"

func sendTickers(tickers string) {

	// Send HTTP POST request
	resp, err := http.Post(tickerDbUrl, "text/plain", bytes.NewBufferString(tickers))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.Status)
		return
	}
}
