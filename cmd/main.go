package main

import (
	"fmt"
	p "golang_webapp/domains/models/produtos"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
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
