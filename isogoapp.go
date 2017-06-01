package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/isomorphicgo/isogoapp/handlers"
)

<<<<<<< HEAD
var webappRoot string = os.Getenv("ISOGO_APP_ROOT")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, webappRoot+"/templates/index.tmpl", nil)
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
=======
var WebAppRoot string = os.Getenv("ISOGO_APP_ROOT")
>>>>>>> 1f58dae263ab74226133033f55a2a449b05b4683

func main() {

	if WebAppRoot == "" {
		fmt.Println("The ISOGO_APP_ROOT environment variable must be set before the web server instance can be started.")
		os.Exit(1)
	}

	fs := http.FileServer(http.Dir(WebAppRoot + "/static"))
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/js/client.js", handlers.GopherjsScriptHandler)
	r.HandleFunc("/js/client.js.map", handlers.GopherjsScriptMapHandler)
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)

}
