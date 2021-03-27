package interfaces

import (
	"librarian/library/data"
)

type Librarian interface {
	// AddUser adds a new user to the registered users.
	AddUser(name string, username string) error

	// RemoveUser removes an existing user from the registered users.
	RemoveUser(username string) error

	// CheckIn checks a lent book, back into the library.
	CheckIn(title string, author string, username string) error

	// CheckOut checks a book in the library out.
	CheckOut(title string, author string, username string) error

	// InStock checks if a book is currently in stock.
	InStock(title string, author string) (bool, error)

	// Stock gets the information about the stock of a book
	Stock(title string, author string) (data.Book, error)
}

type SeniorLibrarian interface {

	// AddBooks adds new stock of a book to the library.
	AddBooks(title string, author string, stock int)

	// RemoveBooks removes stock of a book from the library.
	RemoveBooks(title string, author string, stock int)
}