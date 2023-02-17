package collector

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/pichik/WebWatcher/src/datacenter"
	"github.com/pichik/WebWatcher/src/misc"
)

func ImportTemplate() {
	tmp, err := template.ParseFiles(misc.TemplateDir + "bait.js")
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	baitTemplate = tmp

	webhookTemplate, err = template.ParseFiles(misc.TemplateDir + "webhookTemplate.txt")
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
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

// type MessageContent struct {
// 	Content string `json:"content"`
// }

// func sendToWebhook(data *datacenter.Data) {
//
// 	if misc.Config.Webhook == "" {
// 		return
// 	}
//
// 	var tpl bytes.Buffer
// 	err := webhookTemplate.Execute(&tpl, data)
//
// 	if err != nil {
// 		misc.ErrorLog.Printf("%s", err)
// 	}
//
// 	content := MessageContent{
// 		Content: tpl.String(),
// 	}
//
// 	body, err := json.Marshal(content)
// 	if err != nil {
// 		misc.ErrorLog.Printf("%s", err)
// 	}
// 	misc.DebugLog.Printf("%s", content)
//
// 	// create http client
// 	client := &http.Client{}
//
// 	// create request
// 	req, err := http.NewRequest("POST", misc.Config.Webhook, bytes.NewBuffer(body))
// 	if err != nil {
// 		misc.ErrorLog.Printf("%s", err)
// 	}
//
// 	// set request headers
// 	req.Header.Set("Content-Type", "application/json")
//
// 	// send request
// 	_, err = client.Do(req)
// 	if err != nil {
// 		misc.ErrorLog.Printf("%s", err)
// 	}
// }
