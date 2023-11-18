package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
)

type IndexInput struct {
	Films []film
}

type film struct {
	Name     string
	Director string
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	t := template.Must(template.ParseFiles("index.html"))

	r.Post("/films", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Req: " + r.Header.Get("HX-Request"))
		name := r.PostFormValue("name")
		title := r.PostFormValue("director")
		films = append(films, film{Name: name, Director: title})
		input := IndexInput{Films: films}
		err := t.Execute(w, input)
		if err != nil {
			log.Printf("ERROR: " + err.Error() + "\n")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		input := IndexInput{Films: films}
		err := t.Execute(w, input)
		if err != nil {
			log.Printf("ERROR: " + err.Error() + "\n")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", r)
}

var films = []film{
	{Name: "The Godfather", Director: "Francis Ford Coppola"},
	{Name: "The Shawshank Redemption", Director: "Frank Darabont"},
	{Name: "Schindler's List", Director: "Steven Spielberg"},
}
