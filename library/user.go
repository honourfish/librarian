package library

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the library user who wants to borrow books.
type User struct {
	// ID is the unique mongodb id.
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Name of the user.
	Name string `bson:"title,omitempty"`

	// Books are the books the user has checked out
        Books [] primitive.ObjectID `bson:"books,omitempty'`
}
