package corpus_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/sidpatel93/pocket-sized-go-projects/project_4/gordle/corpus"
)

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		file   string
		length int
		err    error
	}{
		"English corpus": {
			file:   "../corpus/english.txt",
			length: 34,
			err:    nil,
		},
		"empty corpus": {
			file:   "../corpus/empty.txt",
			length: 0,
			err:    corpus.ErrCorpusIsEmpty,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			words, err := corpus.ReadCorpus(tc.file)
			if !errors.Is(tc.err, err) {
				t.Errorf("expected err %v, got %v", tc.err, err)
			}

			if tc.length != len(words) {
				t.Errorf("expected %d, got %d", tc.length, len(words))
			}
		})
	}
}

func TestPickWord(t *testing.T) {
	corpusData := []string{"apple", "banana", "cherry", "date", "elderberry"}

	word := corpus.PickWord(corpusData)
	if !slices.Contains(corpusData, word) {
		t.Errorf("picked word %s not in corpus", word)
	}

}
