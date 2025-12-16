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

const solutionLength = 5

var errInvalidWordLength = fmt.Errorf("input length does not match solution length")

func (e *Game) validateWord(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w",
			solutionLength, len(guess), errInvalidWordLength)
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
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)
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
			return guess
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
