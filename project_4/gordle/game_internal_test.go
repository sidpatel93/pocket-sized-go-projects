package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			gameInstance := New(strings.NewReader(tc.input), "", 0)

			got := gameInstance.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %q, want %q", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateWord(t *testing.T) {
	tt := map[string]struct {
		word      []rune
		expextErr error
	}{
		"valid length": {
			word:      []rune("WORLD"),
			expextErr: nil,
		},
		"invalid length - too short": {
			word:      []rune("HEY"),
			expextErr: errInvalidWordLength,
		},
		"invalid length - too long": {
			word:      []rune("GORDLE"),
			expextErr: errInvalidWordLength,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(nil, "", 0)
			err := g.validateWord(tc.word)
			if !errors.Is(err, tc.expextErr) {
				t.Errorf("%c, expected %q but got %q", tc.word, tc.expextErr, err)
			}
		})
	}
}
