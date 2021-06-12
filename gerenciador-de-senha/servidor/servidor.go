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

//Inserir, insere um servico no banco de dados
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

	statement, erro := db.Prepare("INSERT INTO servico(service,user,password) VALUES (?,?,?)")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
	}
	defer statement.Close()

	insercao, erro := statement.Exec(service.Servico, service.User, service.Password)
	_ = insercao
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	fmt.Println("Usuario inserido.", service.Servico)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario e senha guardados com sucesso, servico: %s", service.Servico)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados."))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT service FROM servico;")
	if erro != nil {
		w.Write([]byte("Erro ao consultar linhas."))
		return
	}
	defer linhas.Close()

	var services []service
	for linhas.Next() {
		var servico service

		if erro := linhas.Scan(&servico.Servico); erro != nil {
			w.Write([]byte("Erro ao scanear o serviço."))
			return
		}
		services = append(services, servico)
	}
	w.WriteHeader(http.StatusOK)
	if erro = json.NewEncoder(w).Encode(services); erro != nil {
		w.Write([]byte("Erro ao converter serviço para json"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	//parametros := mux.Vars(r)

	//servico:= parametros["id"]

}
