package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`
	 /" |  | "\     /""\    (\"   \|"  \  /" _   "| |"  \    /"  |     /""\    (\"   \|"  \  
	(:  (__)  :)   /    \   |.\\   \    |(: ( \___)  \   \  //   |    /    \   |.\\   \    | 
	 \/      \/   /' /\  \  |: \.   \\  | \/ \       /\\  \/.    |   /' /\  \  |: \.   \\  | 
	 //  __  \\  //  __'  \ |.  \    \. | //  \ ___ |: \.        |  //  __'  \ |.  \    \. | 
	(:  (  )  :)/   /  \\  \|    \    \ |(:   _(  _||.  \    /:  | /   /  \\  \|    \    \ | 
	 \__|  |__/(___/    \___)\___|\____\) \_______) |___|\__/|___|(___/    \___)\___|\____\) 
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(left int) {
	var draw string
	switch left {
	case 0:
		draw = `
		  +---+
		  |   |
		  O   |
		 /|\  |
		 / \  |
		      |
		=========`
	case 1:
		draw = `
		  +---+
		  |   |
		  O   |
		 /|\  |
		 /    |
		      |
		=========`
	case 2:
		draw = `
		  +---+
		  |   |
		  O   |
		 /|\  |
		      |
		      |
		=========`
	case 3:
		draw = `
		  +---+
		  |   |
		  O   |
		 /|   |
		      |
		      |
		=========`
	case 4:
		draw = `
		  +---+
		  |   |
		  O   |
		  |   |
		      |
		      |
		=========`
	case 5:
		draw = `
		  +---+
		  |   |
		  O   |
		      |
		      |
		      |
		=========`
	case 6:
		draw = `
		  +---+
		  |   |
		      |
		      |
		      |
		      |
		=========`
	case 7:
		draw = `
		      +
		      |
		      |
		      |
		      |
		      |
		=========`
	case 8:
		draw = `
		`
	}

	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good Guess")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used", guess)
	case "badGuess":
		fmt.Printf("Bad guessn '%s' is not in the word", guess)
	case "lost":
		fmt.Print("You lost! :( The word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WIN! :) The word was: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Print(c, " ")
	}
	fmt.Println()
}
