package collector

import (
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/pichik/server/src/auth"
	"github.com/pichik/server/src/datacenter"
	"github.com/pichik/server/src/misc"
)

var baitTemplate *template.Template

func Baited(w http.ResponseWriter, r *http.Request) bool {

	if auth.IsAuthed(r) || r.URL.Path == datacenter.DeepCollectorPath || strings.HasPrefix(r.URL.Path, "/"+misc.Config.CollectorPath) {
		return false
	}

	if (r.Referer() != "" && !strings.HasPrefix(r.Referer(), misc.Config.Host)) || r.Header.Get("Origin") != "" || r.Header.Get("Accept") == "*/*" || r.Header.Get("Sec-Fetch-Dest") == "script" {
		throwBait(w)
		misc.DebugLog.Printf("[Sending Bait] [%s]%s", r.Method, r.RequestURI)
		return true
	}
	return false
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
	datacenter.Collection.DeepData = append([]*datacenter.Data{data}, datacenter.Collection.DeepData...)

	webhookSend(data)
}

func RequestCollect(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Request Collecting] [%s]%s", r.Method, r.RequestURI)

	var data = datacenter.Data{
		IP:       strings.Split(r.RemoteAddr, ":")[0],
		Method:   r.Method,
		Origin:   r.Header.Get("Origin"),
		Referrer: r.Header.Get("Referer"),
		Location: r.URL.RequestURI(),
		HASH:     datacenter.CreateHash("request"),

		BrowserTime: time.Now().Format("02.01.2006 | 15:04:05"),
	}
	datacenter.Collection.Request = append([]*datacenter.Data{&data}, datacenter.Collection.Request...)

	if strings.HasPrefix(r.URL.Path, "/"+misc.Config.CollectorPath) {
		misc.DebugLog.Printf("[Checking extension] [%s]", r.URL.Path)
		extensionUpgrade(r.URL.Path, w)
	}
}
