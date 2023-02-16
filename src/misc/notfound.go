package misc

import (
	"math/rand"
	"net/http"
	"text/template"
)

var notfoundTemplate *template.Template

type RandomString struct {
	Start string
	End   string
}

func ImportTemplates() {
	tmp, err := template.ParseFiles(TemplateDir + "notFound.html")
	if err != nil {
		ErrorLog.Printf("%s", err)
	}
	notfoundTemplate = tmp
}

func Custom404Handler(w http.ResponseWriter) {
	err := notfoundTemplate.Execute(w, RandomString{Start: RandString(rand.Intn(1000)), End: RandString(rand.Intn(1000))})
	if err != nil {
		ErrorLog.Printf("%s", err)
	}
	w.WriteHeader(http.StatusOK)
}
