package corpus

import (
	"os"
	"strings"
)

func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, ErrReadingCorpus
	}
	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}
	words := strings.Fields(string(data))
	return words, nil
}
