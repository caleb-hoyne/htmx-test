package main

import (
	"github.com/caleb-hoyne/htmx-test/templates"
	"github.com/caleb-hoyne/htmx-test/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/films", func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")
		title := r.PostFormValue("director")
		films = append(films, types.Movie{Name: name, Director: title})
		err := templates.Movies(films).Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := templates.Page(films).Render(ctx, w)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", r)
}

var films = []types.Movie{
	{Name: "The Godfather", Director: "Francis Ford Coppola"},
	{Name: "The Shawshank Redemption", Director: "Frank Darabont"},
	{Name: "Schindler's List", Director: "Steven Spielberg"},
}
