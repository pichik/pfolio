package harvester

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pichik/webwatcher/src/datacenter"
	"github.com/pichik/webwatcher/src/misc"
)

func Extract(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Exctracting] [%s]%s", r.Method, r.RequestURI)

	id := mux.Vars(r)["id"]

	var data *datacenter.Data
	for _, d := range datacenter.GetCollection() {
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
}

func ExtractAll(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Extracting all] [%s]%s", r.Method, r.RequestURI)

	err := harvestListTemplate.Execute(w, datacenter.GetCollection())
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	misc.DebugLog.Printf("[Deleting] [%s]%s", r.Method, r.RequestURI)

	id := mux.Vars(r)["id"]
	datacenter.RemoveFromCollection(id)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	datacenter.ClearCollection()
}
