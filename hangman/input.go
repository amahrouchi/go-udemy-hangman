package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	for {
		// Get the user input
		fmt.Print("What is your letter? ")
		guess, err = reader.ReadString('\n') // Simple quto because the function expects a byte (a single char) not a string
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)

		// Check input length
		if len(guess) != 1 {
			fmt.Printf("Invalid letter input. Letter=%v, length=%v\n", guess, len(guess))
			continue
		}

		// Everything is fine, let's get out of the input loop
		break
	}

	return guess, nil
}
