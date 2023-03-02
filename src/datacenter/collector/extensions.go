package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/pichik/webwatcher/src/misc"
)

type ExtensionData struct {
	Extension   []string `json:"Extension"`
	ContentType string   `json:"Content-Type"`
	Payload     string   `json:"payload"`
}

var extensions = map[string]ExtensionData{}

func ImportExtensions() {
	file, err := ioutil.ReadFile(misc.AssetsDir + "extensions.json")

	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}

	err = json.Unmarshal(file, &extensions)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	misc.DebugLog.Printf("[Extensions available] [%s]", extensions)
}

func extensionUpgrade(endpoint string, w http.ResponseWriter) {

	for _, v := range extensions {
		for _, ext := range v.Extension {
			if path.Ext(endpoint) == ext {
				w.Header().Set("Content-Type", fmt.Sprintf("%s", v.ContentType))
				fmt.Fprintln(w, v.Payload)
				misc.DebugLog.Printf("[Upgrading Extension] [%s]", endpoint)
				return
			}
		}
	}
}
