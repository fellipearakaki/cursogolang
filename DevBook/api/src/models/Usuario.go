package models

import (
	"errors"
	"strings"
)

type Usuario struct {
	ID   uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	CPF  string `json:"cpf,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if etapa == "cadastro" && usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if etapa == "cadastro" && usuario.CPF == "" {
		return errors.New("o campo CPF é obrigatório")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.CPF = strings.TrimSpace(usuario.CPF)
}
