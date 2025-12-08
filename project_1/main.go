package main

import (
	"fmt"
	"flag"
)

type language string

var greetingMap = map[language]string{
	"en": "Hello, World!",
	"es": "¡Hola, Mundo!",
	"fr": "Bonjour le monde!",
	"de": "Hallo, Welt!",
	"it": "Ciao, Mondo!",
}

func main() {
	var userLanguage string
	flag.StringVar(&userLanguage, "lang", "en", "Language code for the greeting (e.g., en, es, fr, de, it)")
	flag.Parse()
	fmt.Println(greeting(language(userLanguage)))
}

func greeting(lang language) string {
	if greeting, exists := greetingMap[lang]; exists {
		return greeting
	}
	return fmt.Sprintf("Greeting not available in the specified language. %q", lang)
}