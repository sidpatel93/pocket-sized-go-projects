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

func findCommonBooks(bookworms []Bookworm) []Book {
	counter := bookcount(bookworms)
	var commonBooks []Book
	for book, count := range counter {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return commonBooks
}

func bookcount(bookworms []Bookworm) map[Book]uint {
	counter := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			counter[book]++
		}
	}
	return counter
}
