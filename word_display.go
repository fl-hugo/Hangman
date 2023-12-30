package Hangman

import "math/rand"

func WordDisplay(Game *HangmanGame) {
	indexes := []int{}
	rune_word := []rune(Game.Word)
	n := len(Game.Word)/2 - 1

	//Choose a random letter
	for i := 0; i < n; i++ {
		letter_index := rand.Intn(len(rune_word))

		//Verify if the index is not already choose
		valid_index := true
		for j := 0; j < len(indexes); j++ {
			if indexes[j] == letter_index {
				valid_index = false
				break
			}
		}
		//Manage the index
		if valid_index {
			indexes = append(indexes, letter_index)
		} else {
			i--
		}
	}

	//Create the word that will be displayed
	rune_word_displayed := []rune{}
	for i := 0; i < len(rune_word); i++ {
		rune_word_displayed = append(rune_word_displayed, '_')
	}
	for i := 0; i < len(indexes); i++ {
		rune_word_displayed[indexes[i]] = rune_word[indexes[i]]
	}

	//Implement the word in the game
	Game.Word_displayed = string(rune_word_displayed)
}
