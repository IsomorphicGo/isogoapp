package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Template rendering function
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, os.Getenv("ISOGO_APP_ROOT")+"/templates/index.tmpl", nil)
}
