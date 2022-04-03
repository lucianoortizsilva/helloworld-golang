package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexao Mysql
)

// Conectar e abrir conexao com Banco de dados
func Conectar() (*sql.DB, error) {

	urlConexao := "ortiz:admin@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
