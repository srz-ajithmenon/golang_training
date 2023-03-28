package main

import "fmt"

type bots interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	english := englishBot{}
	spanish := spanishBot{}

	printGreeting(english)
	printGreeting(spanish)

}

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bots) {
	fmt.Println(b.getGreeting())
}
