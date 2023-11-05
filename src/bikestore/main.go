package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello Bikestore")

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func h1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Bikestore\n")
	io.WriteString(w, r.Method)
}
