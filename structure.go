package Hangman

type HangmanGame struct {
	Word           string
	Word_displayed string
	Letter         string
	Letters_used   []string
	Attempts       int
	Error		   string
	GameComplete   bool
}
