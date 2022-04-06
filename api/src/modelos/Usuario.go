package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
	if erro := u.formatar(etapa); erro != nil {
		return erro
	}
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

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	if u.Senha == "" && etapa == "cadastro" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(u.Senha)
		if erro != nil {
			return erro
		}
		u.Senha = string(senhaComHash)
	}
	return nil
}
