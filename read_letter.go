package Hangman

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func AnalyseLetter(Game *HangmanGame) {
	word := Game.Word
	word_displayed := Game.Word_displayed
	letter_found := []int{}

	//Find the letters in the word
	for i := 0; i < len(word); i++ {
		if string(word[i]) == Game.Letter && string(word_displayed[i]) != Game.Letter {
			letter_found = append(letter_found, i)
		}
	}

	//The letter isn't present in the word
	if len(letter_found) == 0 {
		Game.Error = "This letter is not present in the word."
		Game.Attempts--
	}

	//The letter is present in the word
	if len(letter_found) != 0 {
		rune_word_displayed := []rune(Game.Word_displayed)

		//Display all letters found
		if len(letter_found) == 1 {
			rune_word_displayed[letter_found[0]] = rune_word[letter_found[0]]
		} else {
			for i := 0; i < len(letter_found); i++ {
				index := letter_found[i]
				rune_word_displayed[index] = rune_word[index]
			}
		}
		Game.Word_displayed = string(rune_word_displayed)
	}
}

func ValidLetter(Game *HangmanGame, letter string) {
	for {
		//Variables
		rune_letter := []rune(letter)
		valid_letter := true

		//If the user input is empty
		if len(rune_letter) == 0 {
			Game.Error = "You have not entered any characters."
			valid_letter = false
		}

		//If the text entered by the user isn't a letter
		if len(rune_letter) != 0 && !unicode.IsLetter(rune_letter[0]) || len(rune_letter) > 1 {
			Game.Error = "The character you have entered is not a letter. Please enter a valid character."
			valid_letter = false
		}
 
		//If the letter has already been proposed by the user
		for j := 0; j < len(Game.Letters_used); j++ {
			if Game.Letters_used[j] == letter {
				Game.Error = "You have already proposed this letter, please make another choice."
				valid_letter = false
			}
		}

		//Manage the letter
		if valid_letter {
			Game.Letters_used = append(Game.Letters_used, letter)
			Game.Letter = letter
			break
		}
	}
}
