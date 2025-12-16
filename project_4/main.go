package main

import (
	"os"

	"github.com/sidpatel93/pocket-sized-go-projects/project_4/gordle"
)

func main() {
	gordleGame := gordle.New(os.Stdin, "HELLO", 5)
	gordleGame.Play()
}
