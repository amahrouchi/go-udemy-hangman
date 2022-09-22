package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "b"

	hasLetter := letterInWord(guess, word)

	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "a"

	hasLetter := letterInWord(guess, word)

	if hasLetter {
		t.Errorf("Word %s does not contains letter %s. got=%v", word, guess, hasLetter)
	}
}
