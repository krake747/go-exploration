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

func main() {
	fmt.Println("Hello Bikestore")

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
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
