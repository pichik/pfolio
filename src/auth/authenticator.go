package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/pichik/pfolio/src/data"
)

var AdminPanel = "/admin"

func CanAccess(r *http.Request) bool {
	if strings.HasPrefix(r.URL.Path, AdminPanel) && !IsAuthed(r) {
		return false
	}
	return true
}

func IsAuthed(r *http.Request) bool {
	token, err := r.Cookie("token")
	if err != nil {
		return false
	}
	if token.Value != data.Config.Token {
		return false
	}
	return true
}

func Authenticate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Query().Get("token") == data.Config.Token {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    data.Config.Token,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(365 * 24 * time.Hour),
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, AdminPanel+"/all", http.StatusFound)
		return
	}
	next(w, r)
}
