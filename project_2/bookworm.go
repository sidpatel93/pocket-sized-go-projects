package main

import (
	"encoding/json"
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
	var bookworms []Bookworm
	error = json.NewDecoder(file).Decode(&bookworms)
	if error != nil {
		return nil, error
	}
	return bookworms, nil
}
