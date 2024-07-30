package main

import (
	"fmt"
	"myproject/controller"
)

func main() {

	// Display the welcome message
	fmt.Println("==============================================")
	fmt.Println("        Welcome to the Library Management System")
	fmt.Println("==============================================")
	fmt.Println("Please select an option:")
	fmt.Println()

	var choice int

	// Continuously show the menu until the user decides to exit
	for {
		fmt.Println("----------------------------------------------")
		fmt.Println("                  Menu Options                ")
		fmt.Println("----------------------------------------------")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Remove Book")
		fmt.Println("4. Borrow Book")
		fmt.Println("5. Return Book")
		fmt.Println("6. List Available Books")
		fmt.Println("7. List Borrowed Books")
		fmt.Println("8. Exit")
		fmt.Println("----------------------------------------------")
		fmt.Print("Enter your choice (1-8): ")
		fmt.Scanln(&choice)
		fmt.Println()

		switch choice {
		case 1:
			controller.AddBook()
		case 2:
			controller.AddMember()
		case 3:
			controller.RemoveBook()
		case 4:
			controller.BorrowBook()
		case 5:
			controller.ReturnBook()
		case 6:
			controller.ListAvailableBooks()
		case 7:
			controller.ListBorrowedBooks()
		case 8:
			fmt.Println("Exiting the program...")
			fmt.Println("Thank you for using the Library Management System!")
			
			fmt.Println()
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
		fmt.Println()
	}
}
