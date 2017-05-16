package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var webappRoot string = os.Getenv("ISOGO_APP_ROOT")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, webappRoot+"/templates/index.html", nil)
}

// Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func gopherjsScriptHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, webappRoot+"/client/client.js")
}

func gopherjsScriptMapHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, webappRoot+"/client/client.js.map")
}

func main() {

	if webappRoot == "" {
		fmt.Println("The ISOGO_APP_ROOT environment variable must be set before the web server instance can be started.")
		os.Exit(1)
	}

	fs := http.FileServer(http.Dir(webappRoot + "/static"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/js/client.js", gopherjsScriptHandler)
	r.HandleFunc("/js/client.js.map", gopherjsScriptMapHandler)
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)

}
