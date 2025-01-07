package main

import "fmt"

func main() {
	fmt.Println(Hello("Sam", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix

	case "French":
		prefix = frenchHelloPrefix
	}
	return prefix + name
}
