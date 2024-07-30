# Library Management System

![Library Management System](#) <!-- Insert a cover image or logo here -->

## Overview

The **Library Management System** is a console-based application designed to manage a library's inventory of books and its members. It provides an easy-to-use interface for library staff to add and remove books, register new members, and handle book borrowings and returns.

The system is implemented in Go, leveraging simple command-line interactions to perform various library management tasks.

## Features

- **Add Book:** Register new books with details such as title, author, and availability status.
- **Add Member:** Add new members to the library system with unique identifiers.
- **Remove Book:** Remove books from the library collection using their IDs.
- **Borrow Book:** Enable members to borrow available books and track borrowings.
- **Return Book:** Facilitate the return of borrowed books and update their status.
- **List Available Books:** View all books that are available for borrowing.
- **List Borrowed Books:** Display books currently borrowed by a specific member.
- **Exit:** Gracefully exit the application.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) 1.18 or later installed on your machine.
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) for cloning the repository.

### Setup Instructions

Follow these steps to set up and run the Library Management System on your local machine:

1. **Clone the Repository:**

   Open your terminal and execute:

   ```bash
   git clone https://github.com/yourusername/library-management-system.git

   ```

   ```bash
   cd library-management-system
   ```

   ```bash
   go build

   ```

   ```bash
   ./library-management-system
   ```

```

Upon starting the application, you'll be presented with the main menu, offering various options for managing books and members. The interface will look like this:

```

==============================================
Welcome to the Library Management System
==============================================
Please select an option:

---

                  Menu Options

---

1. Add Book
2. Add Member
3. Remove Book
4. Borrow Book
5. Return Book
6. List Available Books
7. List Borrowed Books
8. Exit

---

Enter your choice (1-8):

```


```
