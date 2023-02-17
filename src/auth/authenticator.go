package auth

import (
	"net/http"
	"strings"

	"github.com/pichik/webwatcher/src/misc"
)

var AdminPanel = "/results/"

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
	if token.Value != misc.Config.Token {
		return false
	}
	return true
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("token") == misc.Config.Token {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    misc.Config.Token,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, AdminPanel+"all", http.StatusFound)
		return
	}
}
