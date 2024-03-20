package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pichik/pfolio/src/data"
	"github.com/pichik/pfolio/src/database"
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

	xtbcsv, err := data.ProcessCSVFile(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing csv file: %v", err), http.StatusInternalServerError)
		return
	}
	xtb_usdJSON, err := json.Marshal(xtbcsv)
	if err != nil {
		fmt.Println("Error marshaling XTB_USD to JSON:", err)
		return
	}
	username, err := r.Cookie("username")
	if err != nil {
		fmt.Println("Cookie not found:", err)
		return
	}
	database.AddPfolio(username.Value, xtb_usdJSON, currency)

	w.WriteHeader(http.StatusOK)
}

func getTimeStamp(timeStr string) int64 {
	// Define the layout for the input string
	layout := "02.01.2006 15:04:05"

	// Parse the input string into a time.Time object
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}

	// Convert the time.Time object to a Unix timestamp (seconds since January 1, 1970 UTC)
	timestamp := t.Unix()

	return timestamp
}
