package main

import (
	"log"
	"os"

	private "github.com/breno-ca/private-repo"
)

func main() {
	out := private.PrivateRepoType{}

	out.PrivateMethod()

	log.Println(os.Getenv("TEST_TOKEN"))
}
