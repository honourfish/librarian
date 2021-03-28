package main

import (
	"github.com/stretchr/testify/assert"

	"librarian/client/client/operations"
	"librarian/client/models"

	"librarian/library"
	"librarian/library/data"
)

var Librarian library.Librarian
var User *data.User

func aLibrarianWithUsername(role, username string) error {
	// for now assume one exists

	// only need username for now
	Librarian.Username = username
	
	return nil
}

func aBookIsAddedWithTitleAuthorAndCopies(title, author string, copies int) error {
	// use default constuctor to set the default request timeout
	params := operations.NewPostLibrarianUsernameBookParams()

	Book = &models.Book {
		Title: title,
		Author: author,
		Copies: int64(copies),
	}

	params.Username = Librarian.Username
	params.Book = Book

	if _, err := HttpClient.Operations.PostLibrarianUsernameBook(params); err != nil {
		return err
	}

	return nil
}

func moreCopiesOfTheBookAreAdded(copies int) error {
	// use default constuctor to set the default request timeout
	params := operations.NewPostLibrarianUsernameBookParams()

	Book.Copies = int64(copies)

	params.Username = Librarian.Username
	params.Book = Book

	if _, err := HttpClient.Operations.PostLibrarianUsernameBook(params); err != nil {
		return err
	}

	return nil
}

func copiesOfTheBookAreRemoved(copies int) error {
	// use default constuctor to set the default request timeout
	params := operations.NewDeleteLibrarianUsernameBookTitleAuthorCopiesParams()

	Book.Copies = int64(copies)

	params.Username = Librarian.Username
	params.Title = Book.Title
	params.Author = Book.Author

	if _, err := HttpClient.Operations.DeleteLibrarianUsernameBookTitleAuthorCopies(params); err != nil {
		return err
	}

	return nil
}

// theBookHasCopies assumes the book has already been retrieved and stored in the global 'Book'.
func theBookHasCopies(copies int) error {
	assert.Equal(&t, int64(copies), Book.Copies, "copies are not equal")
	return nil
}