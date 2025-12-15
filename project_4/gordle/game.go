package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Game represents a single instance of a Gordle game.
type Game struct {
	reader bufio.Reader
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
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: *bufio.NewReader(playerInput),
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
	guess := g.ask()
	fmt.Printf("Your guess is : %s\n", string(guess))
}
