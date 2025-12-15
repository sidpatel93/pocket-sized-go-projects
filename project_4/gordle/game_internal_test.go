package gordle

import (
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
			gameInstance := New(strings.NewReader(tc.input))

			got := gameInstance.ask()
			if string(got) != string(tc.want) {
				t.Errorf("got %q, want %q", string(got), string(tc.want))
			}
		})
	}
}
