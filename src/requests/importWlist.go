package request

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pichik/pfolio/src/auth"
)

func ImportWlist(w http.ResponseWriter, r *http.Request) {

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max size
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing multipart form: %v", err), http.StatusInternalServerError)
		return
	}

	tickerStr := r.FormValue("tickers")

	//Check if user is authed to add stock to database
	if auth.IsAuthed(r) {
		sendTickers(tickerStr)
	}

	// Retrieve the form values
	tickers := strings.Split(tickerStr, " ")
	// Retrieve the existing wlist cookie
	wListCookie, err := r.Cookie("wList")

	if err == nil && wListCookie != nil {
		tickers = append(tickers, strings.Split(wListCookie.Value, ",")...)
	}

	tickers = deduplicate(tickers)

	cookieValue := strings.Join(tickers, ",")

	cookie := &http.Cookie{
		Name:   "wList",
		Value:  cookieValue,
		MaxAge: 365 * 24 * 60 * 60, // 1 year in seconds
	}
	// Set the cookie in the response
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func deduplicate(tickers []string) []string {
	// Create a map to store unique elements
	uniqueElements := make(map[string]bool)

	// Add elements from the merged list to the map
	for _, item := range tickers {
		uniqueElements[item] = true
	}

	// Convert the map keys back to a slice
	finalList := make([]string, 0, len(uniqueElements))
	for item := range uniqueElements {
		finalList = append(finalList, item)
	}

	return finalList
}
