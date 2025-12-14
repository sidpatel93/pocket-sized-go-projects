package gordle

import (
	"bufio"
	"io"
)

// Game represents a single instance of a Gordle game.
type Game struct {
	reader bufio.Reader
}

// New creates and returns a new instance of a Gordle game.
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: *bufio.NewReader(playerInput),
	}
	return g
}

// ask reader input until a valid suggestion is made
func (g *Game) ask() {

}

func (g *Game) Play() {

}
