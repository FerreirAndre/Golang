package servidor

import (
	"encoding/json"
	"fmt"
	"gerenciador-de-senha/banco"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type servico struct {
	Service  string `json:"service"`
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

	var service_ servico

	if erro = json.Unmarshal(corpoReq, &service_); erro != nil {
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

	insercao, erro := statement.Exec(service_.Service, service_.User, service_.Password)
	_ = insercao
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	fmt.Println("Usuario inserido.", service_.Service)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario e senha guardados com sucesso, servico: %s", service_.Service)))
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

	var services []servico
	for linhas.Next() {
		var servico servico

		if erro := linhas.Scan(&servico.Service); erro != nil {
			w.Write([]byte("Erro ao scanear o serviço"))
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
	parametros := mux.Vars(r)

	id := parametros["id"]

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	linha, erro := db.Query("SELECT * FROM servico WHERE service = ?", id)
	if erro != nil {
		w.Write([]byte("Erro ao buscar servico"))
		return
	}

	var servico_ servico

	if linha.Next() {
		if erro = linha.Scan(&servico_.Service, &servico_.User, &servico_.Password); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario"))
		}
	}

	if erro = json.NewEncoder(w).Encode(servico_); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para json"))
		return
	}
}

func Editar(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID := parametros["id"]

	corpoReq, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}
	var servico servico

	if erro = json.Unmarshal(corpoReq, &servico); erro != nil {
		w.Write([]byte("Erro ao converter para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE servico SET service = ?, user = ?, password = ? WHERE service = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
	}
	defer statement.Close()
	if _, erro := statement.Exec(servico.Service, servico.User, servico.Password, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar informações"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func Deletar(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id := parametros["id"]

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM servico WHERE service=?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(id); erro != nil {
		w.Write([]byte("Erro ao executar os statement"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
