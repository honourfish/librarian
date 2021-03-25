package library

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book is a library book that can be borrowed
type Book struct {
	// ID is the unique mongodb id.
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Title of the book.
	Title string `bson:"title,omitempty"`

	// Author of the book.
	Author string `bson:"author,omitempty"`

	// PublishDate is the date the book was published.
	PublishDate time.Time `bson:"publish_date,omitempty"`

	// Copies is the number of copies of the book the library has.
	Copies int `bson:"copies,omitempty"`

	// Users is the foreign key for users that currently have the book checked out
	Users []primitive.ObjectID `bson:"users,omitempty"`
}