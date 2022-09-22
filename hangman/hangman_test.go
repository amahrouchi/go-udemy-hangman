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

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")

	if err == nil {
		t.Error("Error should be returned when using an invalid word.")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)
}

func validState(t *testing.T, expectedState string, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("State should be '%s' got '%s'", expectedState, actualState)
		return false
	}

	return true
}
