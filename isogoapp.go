package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/isomorphicgo/isogoapp/common"
	"github.com/isomorphicgo/isogoapp/handlers"
	"github.com/isomorphicgo/isokit"
)

var WebAppRoot string = os.Getenv("ISOGO_APP_ROOT")

func main() {

	if WebAppRoot == "" {
		fmt.Println("The ISOGO_APP_ROOT environment variable must be set before the web server instance can be started.")
		os.Exit(1)
	}

	env := common.Env{}
	ts := isokit.NewTemplateSet()
	ts.GatherTemplates()
	env.TemplateSet = ts
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.IndexHandler)
	r.Handle("/js/client.js", isokit.GopherjsScriptHandler(WebAppRoot))
	r.Handle("/js/client.js.map", isokit.GopherjsScriptMapHandler(WebAppRoot))
	r.Handle("/template-bundle", handlers.TemplateBundleHandler(&env))

	fs := http.FileServer(http.Dir(WebAppRoot + "/static"))
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)

}
