package main

import (
	"database/sql"
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
	http.HandleFunc("/", index)       //Define rota para a função index
	http.ListenAndServe(":8000", nil) //Disponibiliza escuta na porta 8000

}

func index(w http.ResponseWriter, r *http.Request) {

	db := dbConnect() //Abre conexão com db

	selectAllProducts, err := db.Query("Select * from produtos") //Configura variável selectAllProducts para executar select de todos os produtos
	//verifica possíveis erros
	if err != nil {
		panic(err.Error())
	}

	//Estancia a struct de produtos e sua lista
	produto := p.Produto{}
	produtos := []p.Produto{}

	for selectAllProducts.Next() { // O next é um método que busca a próxima linha do resultado da consulta de banco
		var id, quantidade int
		var nome, descricao string
		var preco float64

		//verifica possíveis erros
		err := selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		//Atribui os resultados da linha em questão às variáveis da struct
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		//Apensa na lista de produtos
		produtos = append(produtos, produto)

	}

	temp.ExecuteTemplate(w, "Index", produtos) //o nome deve ser o mesmo definido quando embedou o html, produtos são os dados que serão manipulados no html
	defer db.Close()                           //Fecha conexão com db
}
