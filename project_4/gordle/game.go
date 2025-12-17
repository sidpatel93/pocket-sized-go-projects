package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Game represents a single instance of a Gordle game.
type Game struct {
	reader      bufio.Reader
	solution    []rune
	maxAttempts int
}

var errInvalidWordLength = fmt.Errorf("input length does not match solution length")

func (e *Game) validateWord(guess []rune) error {
	if len(guess) != len(e.solution) {
		return fmt.Errorf("expected %d, got %d, %w",
			len(e.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

// splitToUppercaseCharacters is a naive implementation to turn a string
// into a list of characters.
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// New creates and returns a new instance of a Gordle game.
func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      *bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}
	return g
}

// ask reader input until a valid suggestion is made
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))
	for {
		input, _, err := g.reader.ReadLine()
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}
		guess := []rune(string(input))
		if err := g.validateWord(guess); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
		} else {
			return splitToUppercaseCharacters(string(guess))
		}
	}
}

func (g *Game) Play() {
	fmt.Printf("Welcome to the game of gordle !")
	for currentGuess := 1; currentGuess <= g.maxAttempts; currentGuess++ {
		fmt.Printf("Attempt %d of %d\n", currentGuess, g.maxAttempts)
		guess := g.ask()
		if slices.Equal(guess, g.solution) {
			fmt.Printf("Congratulations! You've guessed the correct word: %s\n", string(g.solution))
			return
		}
	}
}

// computeFeedback verifies every character of the guess against the solution.
func computeFeedback(guess, solution []rune) feedback {
	// initialise holders for marks
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution have different lengths: %d vs %d", len(guess), len(solution))
		// return a feedback full of absent characters
		return result
	}

	// check for correct letters
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctCharacter
			used[posInGuess] = true
		}
	}

	// look for letters in the wrong position
	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			// The character has already been marked, ignore it.
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				// The letter of the solution is already assigned to a letter of the guess.
				// Skip to the next letter of the solution.
				continue
			}

			if character == target {
				result[posInGuess] = wrongPositionCharacter
				used[posInSolution] = true
				// Skip to the next letter of the guess.
				break
			}
		}
	}

	return result
}
