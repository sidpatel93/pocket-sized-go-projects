package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Game represents a single instance of a Gordle game.
type Game struct {
	reader bufio.Reader
}

const solutionLength = 5

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
		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution! Expected %d characters, got %d.\n", solutionLength, len(guess))
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
