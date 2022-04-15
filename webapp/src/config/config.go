package config

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL = ""
	Porta  = 0
	// HaskKey UTILIZADO PARA AUTENTICAR O COOKIE
	HaskKey []byte
	// BlockKey UTILIZADO PARA CRIPTOGRAFAR OS DADOS DO COOKIE
	BlockKey []byte
)

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	if runtime.GOOS == "windows" {
		fmt.Println("\nRodando no windows")
		APIURL = os.Getenv("API_URL_SEM_DOCKER")
	} else {
		fmt.Println("\nRodando no linux")
		APIURL = os.Getenv("API_URL_COM_DOCKER")
	}

	HaskKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
