package harvester

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/pichik/webwatcher/src/auth"
	"github.com/pichik/webwatcher/src/datacenter"
	"github.com/pichik/webwatcher/src/misc"
)

func Extract(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Exctracting] [%s]%s", r.Method, r.RequestURI)

	id := mux.Vars(r)["id"]

	var data *datacenter.Data
	for _, d := range datacenter.GetCollection("0") {
		if d.HASH == id {
			data = d
			break
		}
	}
	if data == nil {
		misc.ErrorLog.Printf("%s - No data found with this id", id)
		return
	}
	data.Screenshot = strings.TrimPrefix(data.Screenshot, "data:image/png;base64,")

	err := harvestDataTemplate.Execute(w, data)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	w.WriteHeader(http.StatusOK)
}

func ExtractAll(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Extracting all] [%s]%s", r.Method, r.RequestURI)

	// if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {

	// } else {
	refreshRate := r.URL.Query().Get("refresh")
	refreshCookie, err := r.Cookie("refresh")

	if refreshRate != "" && (err != nil || refreshCookie.Value != refreshRate) {
		misc.DebugLog.Printf("[Adding refresh cookie] [%s]%s", r.Method, r.RequestURI)
		cookie := &http.Cookie{
			Name:     "refresh",
			Value:    refreshRate,
			HttpOnly: false,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(365 * 24 * time.Hour),
		}
		http.SetCookie(w, cookie)
	}

	err = harvestListTemplate.Execute(w, auth.AdminPanel)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	// }
	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Deleting] [%s]%s", r.Method, r.RequestURI)

	id := mux.Vars(r)["id"]
	datacenter.RemoveFromCollection(id)
	w.WriteHeader(http.StatusOK)
}

// func DeleteAll(w http.ResponseWriter, r *http.Request) {
// 	datacenter.ClearCollection()
// 	w.WriteHeader(http.StatusOK)
// }

func UpdateAll(w http.ResponseWriter, r *http.Request) {

	filterQuery := r.URL.Query().Get("filter")
	filterCookie, err := r.Cookie("filter")

	if filterQuery != "" && (err != nil || filterCookie.Value != filterQuery) {
		misc.DebugLog.Printf("[Adding filter cookie] [%s]%s", r.Method, r.RequestURI)
		cookie := &http.Cookie{
			Name:     "filter",
			Value:    filterQuery,
			HttpOnly: false,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(365 * 24 * time.Hour),
		}
		http.SetCookie(w, cookie)
	}

	var timestamp string
	if filterQuery != "" {
		timestamp = strings.Split(filterQuery, "/")[1]
	} else if err == nil {
		timestamp = strings.Split(filterCookie.Value, "/")[1]
	} else {
		timestamp = "0"
	}

	// If it is an AJAX request, return the data as JSON
	data, err := json.Marshal(datacenter.GetCollection(timestamp))
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	w.WriteHeader(http.StatusOK)
}
