package persister

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"librarian/library"
)

type Persister interface {

	// InitPersister sets the persister up for use.
	InitPersister() error

	// ClosePersister closes the persister.
	ClosePersister() error

}

//type LibrarianPersister interface {
//	AddLibrarian(librarian)
//}

type BookPersister interface {
	// AddBook persists a given book.
	AddBook(book library.Book) error

	// GetBook gets a book from persistent storage.
	GetBook(id primitive.ObjectID) (library.Book, error)

	// DeleteBook delets a persisted book.
	//DeleteBook(id primitive.ObjectID) error

	// UpdateBook updates the persisted book.
	//UpdateBook(id primitive.ObjectID, book librarian.Book) error
}