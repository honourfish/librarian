package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the library user who wants to borrow books.
type User struct {
	// ID is the unique mongodb id.
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Name of the user, non unique.
	Name string `bson:"name,omitempty"`

	// UserName is the unique identifier for the user
	UserName string `bson:"username,omitempty"`

	// CheckedOutBooks are the books the user has checked out
	CheckedOutBooks []CheckedOutBook `bson:"books,omitempty"`
}
