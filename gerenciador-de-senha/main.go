package main

import (
	"fmt"
	"gerenciador-de-senha/servidor"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/service/inserir", servidor.Inserir).Methods(http.MethodPost)
	router.HandleFunc("/service/buscar",servidor.Buscar).Methods(http.MethodGet)
	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
