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
	word := dictionary.PickWord()
	g, err := hangman.New(8, word)
	if err != nil {
		fmt.Printf("The selected word was invalid. word=%s\n", word)
		os.Exit(1)
	}

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
