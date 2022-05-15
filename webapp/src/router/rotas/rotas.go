package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da aplicacao web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...) // (...) pra pegar cada item do slice e ir adicionando
	rotas = append(rotas, rotaPaginaPrincipal)

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	fileServer := http.FileServer(http.Dir("./assets/")) // dizendo pro go onde estao os css e js
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))

	return router
}
