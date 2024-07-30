package controller

import (
	"fmt"
	model "myproject/models"
	services "myproject/services"
)

// Global variables for keeping track of IDs
var (
	id       int
	userid   int
	bookname string
)

// library is a global instance of the Library service
var library = services.NewLibrary()

// AddBook handles adding a new book via user input
func AddBook() {
	var author string
	var status string
	var title string
	fmt.Println("-------------------------------------")
	fmt.Println("        Enter the title of the book        ")
	fmt.Println("-------------------------------------")
	fmt.Scanln(&title)

	fmt.Println("-------------------------------------")
	fmt.Println("        Enter the author of the book        ")
	fmt.Println("-------------------------------------")
	fmt.Scanln(&author)

	fmt.Println("-------------------------------------")
	fmt.Println("        Enter the status of the book        ")
	fmt.Println("-------------------------------------")
	fmt.Scanln(&status)
	id++

	book := model.Book{ID: id, Title: title, Author: author, Status: status}
	library.AddBook(book)
}

// RemoveBook handles removing a book via user input
func RemoveBook() {
	var id int
	fmt.Println("----------------------------------")
	fmt.Println("       Enter the id of the book to remove       ")
	fmt.Println("----------------------------------")
	fmt.Scanln(&id)
	library.RemoveBook(id)
}

// AddMember handles adding a new member via user input
func AddMember() {
	var name string
	fmt.Println("--------------------------------")
	fmt.Println("        Enter the name of the member        ")
	fmt.Println("--------------------------------")
	fmt.Scanln(&name)
	userid++
	member := model.Member{ID: userid, Name: name}
	library.AddMember(member)
}

// BorrowBook handles borrowing a book via user input
func BorrowBook() {
	var id int
	var userid int
	fmt.Println("--------------------------------")
	fmt.Println("        Enter the id of the book to borrow       ")
	fmt.Println("--------------------------------")
	fmt.Scanln(&id)

	fmt.Println("--------------------------------")
	fmt.Println("        Enter the id of the member        ")
	fmt.Println("--------------------------------")
	fmt.Scanln(&userid)

	err := library.BorrowBook(id, userid)

	if err != nil {
		fmt.Println("--------------------------------")
		fmt.Println("               Error                ")
		fmt.Println("--------------------------------")
		fmt.Println(err)
		fmt.Println()
	} else {
		fmt.Println("--------------------------------")
		fmt.Println("        Book borrowed successfully       ")
		fmt.Println("--------------------------------")
		fmt.Println()
	}
}

// ReturnBook handles returning a book via user input
func ReturnBook() {
	var id int
	var userid int
	fmt.Println("--------------------------------")
	fmt.Println("        Enter the id of the book to return        ")
	fmt.Println("--------------------------------")
	fmt.Scanln(&id)

	fmt.Println("--------------------------------")
	fmt.Println("        Enter the id of the member        ")
	fmt.Println("--------------------------------")
	fmt.Scanln(&userid)

	err := library.ReturnBook(id, userid)

	if err != nil {
		fmt.Println("--------------------------------")
		fmt.Println("               Error                ")
		fmt.Println("--------------------------------")
		fmt.Println(err)
		fmt.Println()
	} else {
		fmt.Println("--------------------------------")
		fmt.Println("        Book returned successfully       ")
		fmt.Println("--------------------------------")
		fmt.Println()
	}
}

// ListAvailableBooks lists all available books
func ListAvailableBooks() {
	availableBooks := library.ListAvailableBooks()

	if len(availableBooks) == 0 {
		fmt.Println("--------------------------------")
		fmt.Println("        No books available        ")
		fmt.Println("--------------------------------")
		fmt.Println()
		return
	}

	fmt.Println("------------------------------------------------")
	fmt.Println("      Available Books      ")
	fmt.Println("------------------------------------------------")
	for _, book := range availableBooks {
		fmt.Printf("Title: %s\nAuthor: %s\nStatus: %s\n", book.Title, book.Author, book.Status)
		fmt.Println("------------------------------------------------")
	}
	fmt.Println()
}

// ListBorrowedBooks lists all books borrowed by a specific member
func ListBorrowedBooks() {
	fmt.Println("--------------------------------")
	fmt.Println("        Enter the id of the member        ")
	fmt.Println("--------------------------------")
	var id int
	fmt.Scanln(&id)

	borrowedBooks := library.ListBorrowedBooks(id)
	if len(borrowedBooks) == 0 {
		fmt.Println("--------------------------------")
		fmt.Println("        No books borrowed        ")
		fmt.Println("--------------------------------")
		fmt.Println()
		return
	}

	fmt.Println("------------------------------------------------")
	fmt.Printf("Borrowed Books by Member ID: %d\n", id)
	fmt.Println("------------------------------------------------")
	for _, book := range borrowedBooks {
		fmt.Printf("Title: %s\nAuthor: %s\nStatus: %s\n", book.Title, book.Author, book.Status)
		fmt.Println("------------------------------------------------")
	}
	fmt.Println()
}
