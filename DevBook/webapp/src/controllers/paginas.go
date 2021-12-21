package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func CarregarTeladeLogin(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "login.html", nil)

}

func CarregarTeladeCadastro(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "cadastro.html", nil)

}
