package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type BikeBrand struct {
	Id   int
	Name string
}

func h1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	bikebrands := map[string][]BikeBrand{
		"BikeBrands": {
			{Id: 1, Name: "Trek"},
			{Id: 2, Name: "Giant"},
			{Id: 3, Name: "Orbea"},
		},
	}
	tmpl.Execute(w, bikebrands)

}

func main() {
	fmt.Println("Hello Bikestore")

	fs := http.FileServer(http.Dir("wwwroot"))
	http.Handle("/wwwroot/", http.StripPrefix("/wwwroot/", fs))

	http.HandleFunc("/", h1)

	fmt.Println("Listening on http://127.0.0.1:8000")
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
