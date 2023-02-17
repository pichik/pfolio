package collector

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/pichik/webwatcher/src/datacenter"
	"github.com/pichik/webwatcher/src/misc"
)

func ImportTemplate() {
	tmp, err := template.ParseFiles(misc.TemplateDir + "bait.js")
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	baitTemplate = tmp
}

func throwBait(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/javascript")
	err := baitTemplate.Execute(w, datacenter.DeepCollectorRef)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
}

func extractJson(jsonData []byte) *datacenter.Data {
	var data datacenter.Data

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	return &data
}
