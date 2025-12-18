package main

import (
	"bufio"
	"os"

	"github.com/sidpatel93/pocket-sized-go-projects/project_4/gordle"
	"github.com/sidpatel93/pocket-sized-go-projects/project_4/gordle/corpus"
)

func main() {
	corpusData, err := corpus.ReadCorpus("./gordle/corpus/english.txt")
	if err != nil {
		panic(err)
	}
	gordleGame, err := gordle.New(bufio.NewReader(os.Stdin), corpusData, 5)
	if err != nil {
		panic(err)
	}
	gordleGame.Play()
}
