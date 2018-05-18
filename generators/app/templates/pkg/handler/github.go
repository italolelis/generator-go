package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gojektech/heimdall"
)

// GithubRepos simple handler example
func GithubRepos(client heimdall.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, _ := http.NewRequest(http.MethodGet, "https://api.github.com/users/italolelis/repos", nil)

		res, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, "could not make the request")
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, "could not read github's json")
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(body)
	}
}
