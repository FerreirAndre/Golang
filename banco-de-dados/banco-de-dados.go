package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConn := "andre:devbook@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConn)
	if erro != nil {
		log.Fatal("erro ao conectar")
	}

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão aberta")

	defer db.Close()
}
