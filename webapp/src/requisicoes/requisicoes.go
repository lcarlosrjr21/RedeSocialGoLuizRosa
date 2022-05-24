package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// FazerRequisicaoComAutencicacao é utilizada para colocar o token na requisicao
func FazerRequisicaoComAutencicacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(metodo, url, dados) // cria a requisicao porem ainda não chama - ainda falta o cookie
	if erro != nil {
		return nil, erro
	}

	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}
