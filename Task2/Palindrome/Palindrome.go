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

func isPalindrome(s string) string {
	length := len(s)
	for i := 0; i < length/2; i++ {
		left := strings.ToLower(string(s[i]))
		right := strings.ToLower(string(s[length- i - 1]))
		if left != right {
			return "Not Palindrome"
		}
	}
	return "Palindrome"
}

func main() {

	var word string
	fmt.Println("Enter a word: ")

	// Read the input from the user
	fmt.Scanln(&word)

	// Check if the user entered an empty string
	err := isAlphabetical(word) 

		for err != nil {
		fmt.Print(err)
		fmt.Println(" ,enter a word: ")
		fmt.Scanln(&word)
		err = isAlphabetical(word)
		
	}

	
		// Call the function and pass the user input as an argument
		output := isPalindrome(word)
		
		fmt.Println(output)

}