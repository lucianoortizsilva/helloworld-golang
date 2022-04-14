package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Remove os dados de autenticacao salvos no browser
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
