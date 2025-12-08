package main

import "testing"

func TestGreeting(t *testing.T) {
	type testCase struct {
		lang     string
		expected string
	}
	tests := map[string]testCase{
		"English": {
			lang: "en",
			expected: "Hello, World!",
		},
		"Spanish": {
			lang: "es",
			expected: "¡Hola, Mundo!",
		},
		"French": {
			lang: "fr",
			expected: "Bonjour le monde!",
		},
		"German": {
			lang: "de",
			expected: "Hallo, Welt!",
		},
		"Italian": {
			lang: "it",
			expected: "Ciao, Mondo!",
		},
		"Japanese": {
			lang: "jp",
			expected: `Greeting not available in the specified language. "jp"`,
		},
		"": {
			lang: "",
			expected: `Greeting not available in the specified language. ""`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := greeting(language(tt.lang))
			if result != tt.expected {
				t.Errorf("greeting(%q) = %q; want %q", tt.lang, result, tt.expected)
			}
		})
	}
}
