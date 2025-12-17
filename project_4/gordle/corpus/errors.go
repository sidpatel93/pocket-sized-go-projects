package corpus

type corpusError string

func (e corpusError) Error() string {
	return string(e)
}

const ErrCorpusIsEmpty = corpusError("corpus is empty")
const ErrReadingCorpus = corpusError("error reading corpus")
