package main

import (
	"apiGO/database"
	"apiGO/models"
	"apiGO/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComOBancoDeDados()

	fmt.Println("ola go lang :)")
	routes.HandleRequest()
}
