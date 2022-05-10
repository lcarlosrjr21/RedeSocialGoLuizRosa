package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/* func init() {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil { // gerando "secret" em slice de bytes com 64 numeros aleatorios
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave) // convertendo a chave secret em string
	fmt.Println(stringBase64)
} */

func main() {
	config.Carregar()
	r := router.Gerar()

	// fmt.Println(config.SecretKey)

	fmt.Printf("Rodando API - Escutando na porta %d ", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
