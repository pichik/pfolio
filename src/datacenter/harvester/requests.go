package harvester

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pichik/server/src/datacenter"
	"github.com/pichik/server/src/misc"
)

func Extract(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Exctracting] [%s]%s", r.Method, r.RequestURI)

	id := mux.Vars(r)["id"]

	var data *datacenter.Data
	for _, d := range datacenter.Collection.DeepData {
		if d.HASH == id {
			data = d
			break
		}
	}
	if data == nil {
		return
	}

	err := extractorTemplate.Execute(w, data)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	w.Header().Set("HasContent", "true")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Deleting] [%s]%s", r.Method, r.RequestURI)

	id := mux.Vars(r)["id"]

	if r.URL.Query().Get("c") == "data" {
		datacenter.Collection.DeepData = removeData(datacenter.Collection.DeepData, id)
	} else if r.URL.Query().Get("c") == "request" {
		datacenter.Collection.Request = removeData(datacenter.Collection.Request, id)
	}
}

func ExtractAll(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Extracting all] [%s]%s", r.Method, r.RequestURI)

	err := cacheTemplate.Execute(w, datacenter.Collection)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	w.Header().Set("HasContent", "true")
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	datacenter.Collection = datacenter.DataCollection{}
}
