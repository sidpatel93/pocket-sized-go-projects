package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	bufferedReader := bufio.NewReaderSize(file, 1024*1024)
	var bookworms []Bookworm
	decoder := json.NewDecoder(bufferedReader)
	error = decoder.Decode(&bookworms)
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
	return sortBooksByAuthor(commonBooks)
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

// Sort books by author and then title
func sortBooksByAuthor(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
