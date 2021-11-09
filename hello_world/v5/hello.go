package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a personalised greeting, defaulting to Hello, world if an empty name is passed.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(langauge string) (prefix string) {

	switch langauge {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	//Can just 'return' and it will return prefix since it
	//is specified above but I don't like that
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
