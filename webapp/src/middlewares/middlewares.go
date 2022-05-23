package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger escreve informacoes da requisicao no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica a existencia de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valores, erro := cookies.Ler(r)
		fmt.Println(valores, erro)
		proximaFuncao(w, r)
	}
}
