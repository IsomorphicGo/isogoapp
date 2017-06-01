package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/isomorphicgo/isogoapp/handlers"
)

var WebAppRoot string = os.Getenv("ISOGO_APP_ROOT")

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
