package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	stringConn := ""

	db, erro := sql.Open("mysql", stringConn)
	if erro != nil {
		return nil, erro
	}
	fmt.Println("conexão aberta.")
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
