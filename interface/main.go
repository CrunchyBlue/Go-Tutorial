package main

import "fmt"

type bot interface {
	getGreeting() string
	getAngryGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	printAngryGreeting(eb)
	printAngryGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello my friend!"
}

func (spanishBot) getGreeting() string {
	return "¡Hola mi amigo!"
}

func (englishBot) getAngryGreeting() string {
	return "HELLO MY FRIEND!"
}

func (spanishBot) getAngryGreeting() string {
	return "¡HOLA MI AMIGO!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printAngryGreeting(b bot) {
	fmt.Println(b.getAngryGreeting())
}
