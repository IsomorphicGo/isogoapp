package handlers

import (
	"net/http"
	"os"

	"github.com/isomorphicgo/isogoapp/view"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	view.RenderTemplate(w, os.Getenv("ISOGO_APP_ROOT")+"/templates/index.html", nil)
}
