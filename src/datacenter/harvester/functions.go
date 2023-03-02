package harvester

import (
	"html/template"

	"github.com/pichik/webwatcher/src/misc"
)

var harvestDataTemplate *template.Template
var harvestListTemplate *template.Template

func ImportTemplate() {
	tmp, err := template.ParseFiles(misc.TemplateDir + "harvestedData.html")
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	harvestDataTemplate = tmp

	harvestListTemplate, err = template.ParseFiles(misc.TemplateDir + "harvestedList.html")
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
}
