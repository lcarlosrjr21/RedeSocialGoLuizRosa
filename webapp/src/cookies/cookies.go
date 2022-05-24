//go get github.com/gorilla/securecookie

package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configurar utiliza as variaveis de ambiente para a criacao do SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar registra as informacoes de autenticacao
func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true, // evitar acesso ao cookie pelo lado cliente
	})

	return nil
}

// Ler retorna os valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados") // lendo cookies com dados codificados
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string) // alocando map vazio pra variavel valores
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}
