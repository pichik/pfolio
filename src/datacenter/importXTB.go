package datacenter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pichik/pfolio/src/misc"
)

func ImportXTB(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IMPORTING XTB")

	refreshRate := r.URL.Query().Get("refresh")
	refreshCookie, err := r.Cookie("refresh")
	fmt.Println(refreshRate)

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

	w.WriteHeader(http.StatusOK)
}
