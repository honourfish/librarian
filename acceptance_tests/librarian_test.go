package main

import (
	"github.com/stretchr/testify/assert"

	"librarian/client/client/operations"
	"librarian/client/models"

	"librarian/library"
)

var Librarian library.Librarian
var User *models.User

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
	return t.err
}

func aLibraryWithNoUsers() error {
	// assume the test starts with no users
	return nil
}

func aUserIsAddedWithUsernameAndName(username, name string) error {
	// use default constuctor to set the default request timeout
	params := operations.NewPostLibrarianUsernameUserParams()

	User = &models.User {
		Name: name,
		Username: username,
	}

	params.Username = Librarian.Username
	params.User = User

	if _, err := HttpClient.Operations.PostLibrarianUsernameUser(params); err != nil {
		return err
	}

	return nil
}

func theUserCanBeRetrieved() error {
	// use default constuctor to set the default request timeout
	params := operations.NewGetLibrarianUsernameUserUserParams()

	params.Username = Librarian.Username
	params.User = User.Username

	response, err := HttpClient.Operations.GetLibrarianUsernameUserUser(params)
	if err != nil {
		return err
	}

	assert.Equal(&t, *User, *response.Payload, "expected user and actual user are different")
	return t.err
}