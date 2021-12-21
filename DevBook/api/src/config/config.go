package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String de conexão com o banco

	StringConexao = ""

	//porta da API

	Porta = 0
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Porta = 9000
	}

	StringConexao = fmt.Sprintf("%s;%s@/%s?charset=utf8&parseTime=True&loc=local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
