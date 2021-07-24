package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":8080", r)
}
