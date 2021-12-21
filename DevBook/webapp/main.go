package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Rodando webapp")

	utils.CarregarTemplates()
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":8080", r))
}
