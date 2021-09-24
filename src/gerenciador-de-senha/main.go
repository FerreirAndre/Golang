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
	router.HandleFunc("/service", servidor.Inserir).Methods(http.MethodPost)
	router.HandleFunc("/service",servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/service/{id}",servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/service/{id}",servidor.Editar).Methods(http.MethodPut)
	router.HandleFunc("/service/{id}",servidor.Deletar).Methods(http.MethodDelete)
	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
