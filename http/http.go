package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{"André", "adrecruz15@gmail.com"}
	templates.ExecuteTemplate(w, "index.html", u)
}

func main() {
	templates = template.Must(templates.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Escutando porta 5000.")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
