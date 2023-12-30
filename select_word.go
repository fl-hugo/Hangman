package Hangman

import (
	"bufio"
	"math/rand"
	"os"
)

func SelectWord(file string, Game *HangmanGame) {
	//Reading the file
	readFile, _ := os.Open(file)
	var slice []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		slice = append(slice, fileScanner.Text())
	}
	readFile.Close()

	//Randomly select a word in the slice created
	index := rand.Intn(len(slice))
	Game.Word = slice[index]
}
