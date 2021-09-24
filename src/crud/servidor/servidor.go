package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func criarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler corpo da requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao conectar no bando de dados."))
		return
	}

	statement, erro := db.Prepare("INSERT INTO usuarios(nome, email) VALUES(?,?)")

	if erro != nil {
		w.Write([]byte("erro ao preparar statement"))
	}
	defer statement.Close()
	inserir, erro := statement.Exec(usuario.Nome, usuario.Email)
	_ = inserir
	if erro != nil {
		w.Write([]byte("erro ao executar statement"))
		return
	}

	fmt.Println("usuario ", usuario.Nome, "inserido com sucesso")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("usuario inserido com sucesso, usuario: %s", usuario.Nome)))
}

func buscar(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro !=  nil {
		w.Write([]byte("erro ao pesquisar no banco de dados"))
		return
	}
	defer linhas.Close()

	var users []usuario

	for linhas.Next() {
		var user usuario
		
		if erro := linhas.Scan(&user.Nome); erro != nil {
			w.Write([]byte("erro ao ler linha"))
			return
		}
		users = append(users, user)
	}
}