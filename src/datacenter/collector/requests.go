package collector

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pichik/webwatcher/src/auth"
	"github.com/pichik/webwatcher/src/datacenter"
	"github.com/pichik/webwatcher/src/misc"
)

var baitTemplate *template.Template

func Bait(w http.ResponseWriter, r *http.Request) {

	if auth.IsAuthed(r) || r.URL.Path == datacenter.DeepCollectorPath || strings.HasPrefix(r.URL.Path, "/"+misc.Config.CollectorPath) {
		return
	}

	throwBait(w)
	w.WriteHeader(http.StatusOK)

	misc.DebugLog.Printf("[Sending Bait] [%s]%s", r.Method, r.RequestURI)
}

func DeepCollect(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Deep Collecting] [%s]%s", r.Method, r.RequestURI)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
		return
	}

	data := extractJson(body)
	data.IP = strings.Split(r.RemoteAddr, ":")[0]
	data.HASH = datacenter.CreateHash("data")
	data.BrowserTime = time.Now().Format("02.01.2006 | 15:04:05")
	data.Collection = "Deep"

	datacenter.AddToCollection(data)

	webhookSend(data)
}

func SimpleCollect(w http.ResponseWriter, r *http.Request) {
	if auth.IsAuthed(r) || r.URL.Path == datacenter.DeepCollectorPath {
		return
	}

	misc.DebugLog.Printf("[Request Collecting] [%s]%s", r.Method, r.RequestURI)

	var data = datacenter.Data{
		IP:         strings.Split(r.RemoteAddr, ":")[0],
		Method:     r.Method,
		Origin:     r.Header.Get("Origin"),
		Referrer:   r.Header.Get("Referer"),
		UserAgent:  r.UserAgent(),
		Location:   r.URL.RequestURI(),
		HASH:       datacenter.CreateHash("request"),
		Collection: "Simple",

		BrowserTime: time.Now().Format("02.01.2006 | 15:04:05"),
	}
	datacenter.AddToCollection(&data)

	if strings.HasPrefix(r.URL.Path, "/"+misc.Config.CollectorPath) {
		misc.DebugLog.Printf("[Checking extension] [%s]", r.URL.Path)
		extensionUpgrade(r.URL.Path, w)
	}
}
