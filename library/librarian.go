package library

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"librarian/library/data"
	"librarian/library/interfaces"
)

const (
	Normal string = "Librarian"
	Senior string = "Senior"
)

// Librarian is a library administrator that depending on role can
//   checkin or checkout books.
type Librarian struct {
	// ID is the unique mongodb id.
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Title of the book.
	Name string `bson:"title,omitempty"`

	// Role is the librarians role (use constant Normal or Senior)
	Role string `bson:"role,omitempty"`

	// Persister is the persister used to add users
	//   and books to persistant storage.
	Persister interfaces.Persister
}

// book gets the given book (if any) from persistent storage.
func (l *Librarian) book(title string, author string) (data.Book, error) {
	var book data.Book

	filter := bson.D{{"title", title},{"author", author}}

	// attempt to get the book from persistent storage
	err := l.Persister.Retrieve("books", filter, &book)
	if err != nil {
		return book, err
	}

	return book, err
}

// updateBook updates the old_book with the new_book, based on the old_books
//   title and author.
func (l *Librarian) updateBook(old_book data.Book, new_book data.Book) (err error) {

	filter := bson.D{{"title", old_book.Title},{"author", old_book.Author}}

	if err = l.Persister.Update("books", filter, new_book); err != nil {
		return
	}

	return
}

// addBook adds a given book to persistent storage
func (l *Librarian) addBook(book data.Book) (err error) {
	if err = l.Persister.Create("books", book); err != nil {
		return err
	}
	
	return err
}

// AddBooks adds new stock of a book to the library.
func (l *Librarian) AddBooks(title string, author string, stock int) (err error) {

	// first check we are a senior librarian and have the
	// authorisation needed.
	if l.Role != Senior {
		err = &PermissionDeniedError{Role: l.Role}
		return
	}

	// attempt to get the book to see if it already exists
	var book data.Book
	book, err = l.book(title, author)

	// the book doesn't exist so create it
	if err == mongo.ErrNoDocuments {
		book.Title = title
		book.Author = author
		book.Copies = stock

		if err = l.addBook(book); err != nil {
			return
		}
	}

	// check there are no other errors getting the book
	if err != nil {
		return
	}

	// book exists so update it
	book.Copies += stock
	if err = l.updateBook(book, book); err != nil {
		return
	}

	return
}

// RemoveBooks removes stock of a book from the library.
func (l *Librarian) RemoveBooks(title string, author string, stock int) (err error) {

	// first check we are a senior librarian and have the
	// authorisation needed.
	if l.Role != Senior {
		err = &PermissionDeniedError{Role: l.Role}
		return
	}
	return
}