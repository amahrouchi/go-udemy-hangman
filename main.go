package main

import (
	"fmt"
	"tests/go-udemy-hangman/hangman"
)

func main() {
	fmt.Println("Hangman!")

	g := hangman.New(8, "Golang")

	fmt.Println(g)
}
