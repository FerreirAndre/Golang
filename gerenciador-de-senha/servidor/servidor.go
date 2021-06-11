package servidor

import (
	"encoding/json"
	"fmt"
	"gerenciador-de-senha/banco"
	"io/ioutil"
	"net/http"
)

type service struct {
	Servico  string `json:"servico"`
	User     string `json:"user"`
	Password string `json:"password"`
}

//Inserir a função insere um servico no banco de dados
func Inserir(w http.ResponseWriter, r *http.Request) {
	corpoReq, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler corpo da requisição"))
		return
	}

	var service service

	if erro = json.Unmarshal(corpoReq, &service); erro != nil {
		w.Write([]byte("Erro ao converte usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO service(servico,usuario,senha) VALUES (?,?,?)")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
	}
	defer statement.Close()

	insercao, erro := statement.Exec(service.Servico, service.User, service.Password)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	fmt.Println("Usuario inserido.", &insercao)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario e senha guardados com sucesso, servico: %s", service.Servico)))
}
