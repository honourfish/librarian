package library

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibrarianRole int

const (
	Normal LibrarianRole = iota
	Senior
)

// Librarian is a library administrator that depending on role can
//   checkin or checkout books.
type Librarian struct {
	// ID is the unique mongodb id.
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Title of the book.
	Name string `bson:"title,omitempty"`

	// Role is the librarians role (use constant Normal or Senior)
	Role LibrarianRole `bson:"role,omitempty"`
}