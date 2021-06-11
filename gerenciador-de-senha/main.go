package main

import (
	"gerenciador-de-senha/servidor"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/service", servidor.Inserir).Methods(http.MethodPost)
}
