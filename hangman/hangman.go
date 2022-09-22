package hangman

import "strings"

type Game struct {
	State        string   // Game state
	Letters      []string // Letters in the word to guess
	FoundLetters []string // Good guesses
	UsedLetters  []string // Used letters
	TurnsLeft    int      // Remaining attempt
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	return &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
}
