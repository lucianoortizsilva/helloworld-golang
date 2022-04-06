package seguranca

import "golang.org/x/crypto/bcrypt"

// Recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// Compara uma senha e um has, e retorna se ela sao iguais
func VerificarSenha(senhaComHash string, senhaSemHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaSemHash))
}
