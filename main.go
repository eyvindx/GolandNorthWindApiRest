package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main()  {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request)){
		w.Write([]byte("Hola"))
	}
	http.ListemAndServe(": 3000", r)
}