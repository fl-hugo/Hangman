package Hangman

import (
	"bufio"
	"fmt"
	"os"
)

func File() string {
	var file string
	for {
		//Ask the file
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("If you want to use the default file, enter 1.")
		fmt.Print("Otherwise, enter the name of the file you wish to use : ")
		scanner.Scan()
		file = scanner.Text()
		_, valid_file := os.Stat(file)

		//Default file
		if file == "1" {
			return "words.txt"
		}

		//Verify the existence of the file
		if os.IsNotExist(valid_file) {
			fmt.Println("")
			fmt.Println("The file you have entered does not exist. Please enter another file name.")
			fmt.Println("")
		} else {
			break
		}
	}
	return file
}
