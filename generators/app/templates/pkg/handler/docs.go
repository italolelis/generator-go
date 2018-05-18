package handler

import (
	"net/http"
)

// Docs serves the swagger definition file
func Docs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "docs/specification/api.html")
}
