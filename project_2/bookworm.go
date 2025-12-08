package main

import (
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	file, error := os.Open(filePath)
	if error != nil {
		return nil, error
	}
	defer file.Close()
	return nil, nil
}
