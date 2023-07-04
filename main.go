package main

import (
	"fmt"

	"github.com/barbara-teresa-toledo/go-api-rest/models"
	"github.com/barbara-teresa-toledo/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleRequest()
}
