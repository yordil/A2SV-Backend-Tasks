package main

import (
	"fmt"
	"strings"
	"unicode"
)


//  custom error message handler
type Myerror struct {
	message string
}

// Function to return the error message
func (e *Myerror) Error() string { 
	return e.message 
}


// Function to check if the string contains only alphabetical characters
func isAlphabetical(s string) error {

	if s == "" { 
		// Return an error message if the string is empty
		return &Myerror{message: "Empty string is not allowed"}
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return &Myerror{message: "Only alphabetical characters are allowed"}
		}
	}

	return nil
}

func wordCount(word string)   map[string]int {

	// Create a map to store the frequency of each word
	wordMap := make(map[string]int)

	for _, currentLetter := range word {

		// Check if the current letter is a letter to ignore punctuations
		if unicode.IsLetter(currentLetter) {
			// casting to lower case to ignore case sensitivity
			wordMap[strings.ToLower(string(currentLetter))]++
		}
	}

	return wordMap
}

func main() { 

	var word string
	fmt.Println("Enter a sentences: ")

	// Read the input from the user
	fmt.Scanln(&word)

	// Check if the user entered an empty string
	
	err := isAlphabetical(word)

	for err != nil { 
		fmt.Println(err)
		fmt.Println("Enter a sentences: ")
		fmt.Scanln(&word)
		err = isAlphabetical(word)
	}

	frequency := wordCount(word)

	// Displaying the frequency of each word
	for key, value := range frequency {
		fmt.Printf("%s: %d\n", key, value)
	}
	
}