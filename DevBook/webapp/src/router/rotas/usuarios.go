package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuario = []Rota{
	{
		Uri:       "/criar-usuario",
		Metodo:    http.MethodGet,
		Funcao:    controllers.CarregarTeladeCadastro,
		RequerAut: false,
	},

	{
		Uri:       "/usuarios",
		Metodo:    http.MethodPost,
		Funcao:    controllers.CriarUsuario,
		RequerAut: false,
	},
}
