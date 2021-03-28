package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"librarian/library/errors"
)

// User is the library user who wants to borrow books.
type User struct {
	// ID is the unique mongodb id.
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Name of the user, non unique.
	Name string `bson:"name,omitempty"`

	// Username is the unique identifier for the user
	Username string `bson:"username,omitempty"`

	// CheckedOutBooks are the books the user has checked out
	CheckedOutBooks []*CheckedOutBook `bson:"books,omitempty"`
}

// CheckedOutBook gets the checked out book, given the book.
func (u *User) CheckedOutBook(book *Book) (int, *CheckedOutBook, error) {
    for index, checked_out_book := range u.CheckedOutBooks {
        if book.ID == checked_out_book.BookRef {
            return index, checked_out_book, nil
        }
	}
	
    return 0, nil, &errors.NotCheckedOutError{}
}

// RemoveCheckedOutBook removes a given book from the CheckedOutBooks slice.
func(u *User) RemoveCheckedOutBook(book *Book) error {
	// see if the book is checked out first
	index, _, err := u.CheckedOutBook(book)
	if err != nil {
		return err
	}

	// remove from slice (doesn't maintain order)
	slice_len := len(u.CheckedOutBooks)

	u.CheckedOutBooks[index] = u.CheckedOutBooks[slice_len-1] // Copy last element to index 'index'.
	u.CheckedOutBooks[slice_len-1] = nil   // Erase last element.
	u.CheckedOutBooks = u.CheckedOutBooks[:slice_len-1]   // Truncate slice.

	return nil
}