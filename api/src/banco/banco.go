package banco

import (
	"api/src/config"
	"database/sql"
	"fmt"
	"runtime"

	_ "github.com/go-sql-driver/mysql" // Driver
)

func Conectar() (*sql.DB, error) {

	if runtime.GOOS == "windows" {
		fmt.Println("\nRodando no windows")

		db, erro := sql.Open("mysql", config.StringConexaoBancoSemDocker)
		if erro != nil {
			return nil, erro
		}
		if erro = db.Ping(); erro != nil {
			db.Close()
			return nil, erro
		}
		return db, nil
	} else {
		fmt.Println("\nRodando no linux via docker")

		db, erro := sql.Open("mysql", config.StringConexaoBancoComDocker)
		if erro != nil {
			return nil, erro
		}
		if erro = db.Ping(); erro != nil {
			db.Close()
			return nil, erro
		}
		return db, nil
	}

	return nil, nil
}
