package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:       "/",
		Metodo:    http.MethodPost,
		Funcao:    controllers.CriarUsuario,
		RequerAut: false,
	},
	{
		Uri:       "/usuários",
		Metodo:    http.MethodPost,
		Funcao:    controllers.CriarUsuario,
		RequerAut: false,
	},
	{
		Uri:       "/usuários",
		Metodo:    http.MethodPost,
		Funcao:    controllers.CriarUsuario,
		RequerAut: false,
	},
	{
		Uri:       "/usuários/{usuarioId}",
		Metodo:    http.MethodGet,
		Funcao:    controllers.BuscarUsuarioporId,
		RequerAut: false,
	},
	{
		Uri:       "/usuários/{usuarioId}",
		Metodo:    http.MethodPut,
		Funcao:    controllers.AtualizarUsuario,
		RequerAut: false,
	},
	{
		Uri:       "/usuários/{usuarioId}",
		Metodo:    http.MethodDelete,
		Funcao:    controllers.DeletarUsuario,
		RequerAut: false,
	},
}
