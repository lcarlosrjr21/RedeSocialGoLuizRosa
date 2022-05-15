package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"
)

// CriarUsuario chama a API para cadastrar um usuario no BD
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// nome := r.FormValue("nome")

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		// log.Fatal(erro)
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// fmt.Println(usuario)
	// fmt.Println(bytes.NewBuffer(usuario))

	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, erro := http.Post(url, "/application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		// log.Fatal(erro)
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	// fmt.Println(response.Body)

	if response.StatusCode >= 400 { // 400 pra cima é falha na operacao
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}
