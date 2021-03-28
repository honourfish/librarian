package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"librarian/library/errors"
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

// ContainsUser checks that the book contains a given user in its Users member.
func (b *Book) ContainsUser(user *User) bool {
    for _, user_id := range b.Users {
        if user.ID == user_id {
            return true
        }
    }
    return false
}

// User gets the user id from the Users slice, given the user.
func (b *Book) User(user *User) (int, primitive.ObjectID, error) {
    for index, user_id := range b.Users {
        if user.ID == user_id {
            return index, user_id, nil
        }
	}
	
    return 0, primitive.ObjectID{}, &errors.NotCheckedOutError{}
}

// RemoveUser removes a given user id from the Users slice.
func(b *Book) RemoveUser(user *User) error {
	// see if the user has the book checked out first
	index, _, err := b.User(user)
	if err != nil {
		return err
	}

	// remove from slice (doesn't maintain order)
	slice_len := len(b.Users)
	
	b.Users[index] = b.Users[slice_len-1] // Move indexed element to last.
	b.Users[slice_len-1] = primitive.ObjectID{}   // Erase last element.
	b.Users = b.Users[:slice_len-1]   // Truncate slice.

	return nil
}