package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"

	"github.com/pichik/webwatcher/src/datacenter"
	"github.com/pichik/webwatcher/src/misc"
)

func ImportTemplate() {
	filePath := misc.TemplateDir + "bait.js"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	baitContent = content
}

func throwBait(w http.ResponseWriter) {
	w.Header().Set("Content-Type", fmt.Sprintf("text/javascript"))
	_, err := w.Write(baitContent)
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
