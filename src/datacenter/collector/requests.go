package collector

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pichik/webwatcher/src/datacenter"
	"github.com/pichik/webwatcher/src/misc"
)

var baitContent []byte

func Bait(w http.ResponseWriter, r *http.Request) {
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

func SimpleCollect(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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

	next(w, r)
	w.WriteHeader(http.StatusOK)
}

func GetExtension(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Checking extension] [%s]", r.URL.Path)
	extensionUpgrade(r.URL.Path, w)
	w.WriteHeader(http.StatusOK)
}
