package main

import (
	"path/filepath"
	"reflect"
	"testing"
)

var (
	handmaidsTale = Book{
		Author: "Margaret Atwood", Title: "The Handmaid's Tale",
	}
	oryxAndCrake = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar   = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre     = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms(t *testing.T) {
	testcases := map[string]struct {
		filepath string
		want     []Bookworm
		wantErr  bool
	}{
		"valid file": {
			filepath: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"invalid file": {
			filepath: "testdata/non_existent_file.json",
			want:     nil,
			wantErr:  true,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(filepath.Clean(tc.filepath))
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())
			}

			if err == nil && tc.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())

			}
			// We can use the reflect.DeepEqual function to compare the two slices. It is not recommended for production code, but fine for tests.
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
			// if !equalBookworms(t, got, tc.want) {
			// 	t.Fatalf("different result: got %v, expected %v", got, tc.want)
			// }
		})
	}

}

// equalBookworms is a helper to test the equality of two lists of Bookworms.
func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	if len(bookworms) != len(target) {
		// Early exit!
		return false
	}

	for i := range bookworms {
		// Verify the name of the Bookworm.
		if bookworms[i].Name != target[i].Name {
			return false
		}
		// Verify the content of the collections of Books for each Bookworm.
		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	// Everything is equal!
	return true
}

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()

	if len(books) != len(target) {
		// Early exit!
		return false
	}
	// Verify the content of the collections of Books for each Bookworm.
	for i := range target {
		if target[i] != books[i] {
			return false
		}
	}
	// Everything is equal!
	return true
}
