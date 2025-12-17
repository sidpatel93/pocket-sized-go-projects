package corpus

import (
	"fmt"
	"os"
	"strings"
)

const ErrorReadingCorpus = "error reading corpus file"

func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w",
			path, err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("%s: %q is empty", ErrorReadingCorpus, path)
	}
	words := strings.Fields(string(data))
	return words, nil
}
