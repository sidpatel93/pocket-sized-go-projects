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

func Test_computeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:            "HERTZ",
			solution:         "HERTZ",
			expectedFeedback: feedback{correctCharacter, correctCharacter, correctCharacter, correctCharacter, correctCharacter},
		},
		"double character": {
			guess:            "HELLO",
			solution:         "HELLO",
			expectedFeedback: feedback{correctCharacter, correctCharacter, correctCharacter, correctCharacter, correctCharacter},
		},
		"double character with wrong answer": {
			guess:            "HELLL",
			solution:         "HELLO",
			expectedFeedback: feedback{correctCharacter, correctCharacter, correctCharacter, correctCharacter, absentCharacter},
		},
		"five identical, but only two are there": {
			guess:            "LLLLL",
			solution:         "HELLO",
			expectedFeedback: feedback{absentCharacter, absentCharacter, correctCharacter, correctCharacter, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			guess:            "HLLEO",
			solution:         "HELLO",
			expectedFeedback: feedback{correctCharacter, wrongPositionCharacter, correctCharacter, wrongPositionCharacter, correctCharacter},
		},
		"three identical, but not in the right position (from left to right)": {
			guess:            "HLLLO",
			solution:         "HELLO",
			expectedFeedback: feedback{correctCharacter, absentCharacter, correctCharacter, correctCharacter, correctCharacter},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			guess:            "LLLWW",
			solution:         "HELLO",
			expectedFeedback: feedback{wrongPositionCharacter, absentCharacter, correctCharacter, absentCharacter, absentCharacter},
		},
		"swapped characters": {
			guess:            "HOLLE",
			solution:         "HELLO",
			expectedFeedback: feedback{correctCharacter, wrongPositionCharacter, correctCharacter, correctCharacter, wrongPositionCharacter},
		},
		"absent character": {
			guess:            "HULFO",
			solution:         "HELFO",
			expectedFeedback: feedback{correctCharacter, absentCharacter, correctCharacter, correctCharacter, correctCharacter},
		},
		"absent character and incorrect": {
			guess:            "HULPP",
			solution:         "HELPO",
			expectedFeedback: feedback{correctCharacter, absentCharacter, correctCharacter, correctCharacter, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf("guess: %q, got the wrong feedback, expected %v, got %v", tc.guess, tc.expectedFeedback, fb)
			}
		})
	}
}
