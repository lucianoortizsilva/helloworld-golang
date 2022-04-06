package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar IRÁ VALIDAR E FORMATAR O USUARIO RECEBIDO
func (u *Usuario) Preparar(etapa string) error {
	if erro := u.validar(etapa); erro != nil {
		return erro
	}
	u.formatar()
	return nil
}

func (u *Usuario) validar(etapa string) error {

	if u.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if u.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	if u.Senha == "" && etapa == "cadastro" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
