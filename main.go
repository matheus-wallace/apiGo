package main

import (
	"apiGO/models"
	"apiGO/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{{Nome: "Nome 1", Historia: "Historia 1"}, {Nome: "Nome 2", Historia: "Historia 2"}}

	fmt.Println("ola go lang :)")
	routes.HandleRequest()
}
