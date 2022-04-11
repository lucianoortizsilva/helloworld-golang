package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	sql := "insert into usuarios (nome, nick, email, senha) values(?,?,?,?)"
	statement, erro := repositorio.db.Prepare(sql)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar retorna todos os usuários por filtro
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	sql := "select id, nome, nick, email, criadoEm from usuarios  where nome like ? or nick like ?"
	linhas, erro := repositorio.db.Query(sql, nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var u modelos.Usuario
		if erro = linhas.Scan(
			&u.ID,
			&u.Nome,
			&u.Nick,
			&u.Email,
			&u.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, u)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {

	sql := "select id, nome, nick, email, criadoEm from usuarios  where id = ?"
	linhas, erro := repositorio.db.Query(sql, ID)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {

	sql := "update usuarios set nome=?, nick=?, email=? where id=?"
	statement, erro := repositorio.db.Prepare(sql)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Usuarios) Deletar(ID uint64) error {

	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	sql := "select id, senha from usuarios where email = ?"
	linha, erro := repositorio.db.Query(sql, email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioID uint64, seguidorID uint64) error {
	sql := "insert ignore into seguidores(usuario_id, seguidor_id) values (?,?)"
	statement, erro := repositorio.db.Prepare(sql)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Usuarios) PararDeSeguir(usuarioID uint64, seguidorID uint64) error {
	sql := "delete from seguidores where usuario_id = ? and seguidor_id = ?"
	statement, erro := repositorio.db.Prepare(sql)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(`
			SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
			  FROM usuarios u inner join seguidores s on u.id = s.seguidor_id 
			 WHERE s.usuario_id = ?	
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, usuario)
	}
	return seguidores, nil
}

func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(`
			SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
			  FROM usuarios u inner join seguidores s on u.id = s.usuario_id
			 WHERE s.seguidor_id = ?
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	sql := "select senha from usuarios where id = ?"
	linha, erro := repositorio.db.Query(sql, usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil
}

func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, novaSenha string) error {

	sql := "update usuarios set senha = ? where id = ?"
	statement, erro := repositorio.db.Prepare(sql)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(novaSenha, usuarioID)
	if erro != nil {
		return erro
	}
	return nil
}
