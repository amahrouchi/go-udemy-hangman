package main

import (
	"fmt"
	"os"
	"tests/go-udemy-hangman/dictionary"
	"tests/go-udemy-hangman/hangman"
)

func main() {
	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		os.Exit(1)
	}

	hangman.DrawWelcome()
	g := hangman.New(8, dictionary.PickWord())

	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v\n", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}

}
