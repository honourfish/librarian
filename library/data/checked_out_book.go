package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CheckedOutBook contains information about a checked out book
type CheckedOutBook struct {

	// BookRef is the foreign key reference to the book object.
	BookRef primitive.ObjectID `bson:"book_ref,omitempty"`

	// CheckOutDate is the date the book was checked out.
	CheckOutDate time.Time `bson:"checkout_date,omitempty"`

	// DueDate is the date the book is due to be checked in.
	DueDate time.Time `bson:"due_date,omitempty"`
}