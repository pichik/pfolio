package harvester

import (
	"text/template"

	"github.com/pichik/WebWatcher/src/datacenter"
	"github.com/pichik/WebWatcher/src/misc"
)

var extractorTemplate *template.Template
var cacheTemplate *template.Template

func ImportTemplate() {
	tmp, err := template.ParseFiles(misc.TemplateDir + "harvestedData.html")
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	extractorTemplate = tmp

	cacheTemplate, err = template.ParseFiles(misc.TemplateDir + "harvestedList.html")
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
}

func removeData(slice []*datacenter.Data, hash string) []*datacenter.Data {
	for s, cd := range slice {
		if cd.HASH == hash {
			return append(slice[:s], slice[s+1:]...)
		}
	}
	return slice
}
