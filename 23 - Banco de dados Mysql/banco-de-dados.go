package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	urlConexao := "ortiz:admin@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		fmt.Println("ERRO AO CONECTAR")
		log.Fatal(erro)
	}

	fmt.Println("Conectado no Mysql com sucesso!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
