package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

//Cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Insere um novo usuário no banco de dados
func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO USUARIOS (Nome, CPF) VALUES (?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.CPF)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil

	//return 0, nil
}

func (repositorio usuarios) Buscar(nomeouNick string) ([]models.Usuario, error) {
	nomeouNick = fmt.Sprintf("%%%s%%", nomeouNick)

	linhas, erro := repositorio.db.Query("select id, nome, cpf from usuarios where nome like ?",
		nomeouNick)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			usuario.CPF,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

func (repositorio usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuario set nome = ?, cpf = ? where id = ?",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.CPF, ID); erro != nil {
		return erro
	}

	return nil
}
