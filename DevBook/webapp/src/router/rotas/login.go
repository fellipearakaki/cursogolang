package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasLogin = []Rota{
	{
		Uri:       "/",
		Metodo:    http.MethodGet,
		Funcao:    controllers.CarregarTeladeLogin,
		RequerAut: false,
	},
	{
		Uri:       "/login",
		Metodo:    http.MethodGet,
		Funcao:    controllers.CarregarTeladeLogin,
		RequerAut: false,
	},
}
