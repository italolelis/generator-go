package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

// HelloWorld simple handler example
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "Hello world")
}
