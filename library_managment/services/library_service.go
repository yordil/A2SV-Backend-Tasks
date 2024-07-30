package services

import (
	"fmt"
	model "myproject/models"
)

// MyError is a custom error type for handling errors in the library
type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

// LibraryManager defines the interface for library management operations
type LibraryManager interface {
	AddBook(book model.Book)
	AddMember(mem model.Member)
	RemoveBook(id int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []model.Book
	ListBorrowedBooks(memberID int) []model.Book
}

// Library is the concrete implementation of the LibraryManager interface
type Library struct {
	books          map[int]model.Book
	members        map[int]model.Member
	borrowed_books map[int][]int
}

// NewLibrary initializes and returns a new Library instance
func NewLibrary() *Library {
	return &Library{
		books:          make(map[int]model.Book),
		members:        make(map[int]model.Member),
		borrowed_books: make(map[int][]int),
	}
}

// AddMember adds a new member to the library
func (l *Library) AddMember(mem model.Member) {
	fmt.Println("------------------------------------")
	fmt.Println("         Service called: AddMember         ")
	fmt.Println("------------------------------------")
	fmt.Printf("Member ID: %d\nName: %s\n\n", mem.ID, mem.Name)
	l.members[mem.ID] = mem
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book model.Book) {
	fmt.Println("----------------------------------")
	fmt.Println("         Service called: AddBook         ")
	fmt.Println("----------------------------------")
	fmt.Printf("Book ID: %d\nTitle: %s\nAuthor: %s\nStatus: %s\n\n", book.ID, book.Title, book.Author, book.Status)
	l.books[book.ID] = book
}

// RemoveBook removes a book from the library by its ID
func (l *Library) RemoveBook(id int) {
	fmt.Println("----------------------------------")
	fmt.Println("         Service called: RemoveBook         ")
	fmt.Println("----------------------------------")
	
	if _, exists := l.books[id]; !exists { 
		fmt.Println("Book not found")
		return 
	}else{
		fmt.Printf("Book ID: %d has been removed\n\n", id)
		delete(l.books, id)
		return 
	}
}

// BorrowBook allows a member to borrow a book from the library
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExists := l.books[bookID]
	member, memberExists := l.members[memberID]

	if !bookExists {
		return &MyError{message: "Book not found"}
	} else if !memberExists {
		return &MyError{message: "Member not found"}
	} else if book.Status == "available" {
		book.Status = "borrowed"
		l.books[bookID] = book
		member.BorrowedBooks = append(member.BorrowedBooks, book)
		l.members[memberID] = member
		return nil
	} else {
		return &MyError{message: "Book not available"}
	}
}

// ReturnBook allows a member to return a borrowed book to the library
func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, memberExists := l.members[memberID]
	if !memberExists {
		return &MyError{message: "Member not found"}
	}

	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			// Remove the book from the member's borrowed list
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			l.members[memberID] = member

			// Update the book's status to available
			book := l.books[bookID]
			book.Status = "available"
			l.books[bookID] = book

			return nil
		}
	}

	return &MyError{message: "Book not borrowed by this member"}
}

// ListAvailableBooks lists all available books in the library
func (l *Library) ListAvailableBooks() []model.Book {
	var availableBooks []model.Book

	for _, book := range l.books {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}

	return availableBooks
}

// ListBorrowedBooks lists all books borrowed by a specific member
func (l *Library) ListBorrowedBooks(memberID int) []model.Book {
	member, exists := l.members[memberID]
	if !exists {
		return []model.Book{}
	}

	return member.BorrowedBooks
}
