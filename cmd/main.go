package main

import (
	"database/sql"
	"fmt"
	p "golang_webapp/domains/models/produtos"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func dbConnect() *sql.DB { //Função de criaçao de conexão com o postgres
	connString := "user=user dbname=loja password=password host=localhost sslmode=disable" //Configuração da string de conexão (pode ser feito via env)
	db, err := sql.Open("postgres", connString)                                            //Abre conexão
	if err != nil {                                                                        //Valida erros
		panic(err.Error())
	}
	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := dbConnect()                 //Tenta conectar ao db
	defer db.Close()                  //Fecha conexão
	http.HandleFunc("/", index)       //Define rota para a função index
	http.ListenAndServe(":8000", nil) //Disponibiliza escuta na porta 8000

}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []p.Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{Nome: "Tenis", Descricao: "Confortável", Preco: 89, Quantidade: 3},
		{Nome: "Fone", Descricao: "Muito bom", Preco: 59, Quantidade: 2},
		{Nome: "Produto novo", Descricao: "Muito Legal", Preco: 1.99, Quantidade: 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos) //o nome deve ser o mesmo definido quando embedou o html
	fmt.Println(produtos)
}
