package main

import (
	"fmt"
	"os"
	"tests/go-udemy-hangman/hangman"
)

func main() {
	g := hangman.New(8, "Golang")
	fmt.Println(g)

	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("Could not read from terminal: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(l)
}
