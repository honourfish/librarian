package librarian

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book is a library book that can be borrowed
type Book struct {
	// ID is the unique mongodb id.
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Title of the book
	Title string `bson:"title,omitempty"`

	// Author of the book
	Author string `bson:"author,omitempty"`

	// PublishDate is the date the book was published
	PublishDate time.Time `bson:"publish_date,omitempty"`
}