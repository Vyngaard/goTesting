package main

import "fmt"

//important to refactor for efficiency
//both the production code and the tests
const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

//easy to read function containing our switch for language selection
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func main() {
	fmt.Println(Hello("world", "French"))
}
