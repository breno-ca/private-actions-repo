package main

import (
	"log"
	"os"

	private "github.com/breno-ca/private-repo"
)

func main() {
	out := private.PrivateRepoType{}

	out.PrivateMethod()

	token_message := os.Getenv("TEST_TOKEN")

	expected_message := "token de teste, será que consigo ver executando um cat? será que o código consegue usar pra teste?"

	if token_message == expected_message {
		log.Println("mensagem do token lida com sucesso!")
		return
	}

	log.Println("não foi possível ler a mensagem do token")
}
