package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("serving in 3000")
	r := chi.NewRouter()

	//logging and recoverer
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//route to use the API
	r.Post("/find-pairs", FindPairs)

	//port in which the API is serving in
	http.ListenAndServe(":3000", r)
}
