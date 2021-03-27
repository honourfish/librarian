package main

import (
	"fmt"
	"github.com/cucumber/godog"

	"github.com/stretchr/testify/assert"

	"librarian/client/client"
	"librarian/client/client/operations"
	"librarian/client/models"
)

// asserter is a stubbed substitute for testing.T
type asserter struct {
	err error
}

// Errorf is needed for the asserter to be used instead of testing.T
//  So create a wrapper for fmt.Errorf.
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

// Globals can be shared between steps
var HttpClient *client.Library
var Book *models.Book
var t asserter

func aBookIsAddedWithTitleAndAuthor(title, author string) (err error) {
	// use default constuctor to set the default request timeout
	params := operations.NewPostBookParams()

	Book = &models.Book {
		Title: title,
		Author: author,
	}
	params.Book = Book

	//var response *operations.PostBookCreated
	if _, err = HttpClient.Operations.PostBook(params); err != nil {
		return err
	}

	return nil
}

func aLibraryWithNoBooks() error {
	// nothing in the library yet
	return nil
}

func theBookIsUpdatedWithNewTitle(title string) (err error) {
	// use default constuctor to set the default request timeout
	params := operations.NewPutBookTitleParams()

	// set the search filter title to be the old title
	params.Title = Book.Title

	// update the new title
	Book.Title = title
	params.Book = Book

	//var response *operations.PostBookCreated
	if _, err = HttpClient.Operations.PutBookTitle(params); err != nil {
		return err
	}

	return nil
}

func theBookCanBeRetrieved() error {
	// use default constuctor to set the default request timeout
	params := operations.NewGetBookTitleParams()

	params.Title = Book.Title

	response, err := HttpClient.Operations.GetBookTitle(params)
	if err != nil {
		return err
	}

	assert.Equal(&t, *response.Payload, *Book, "expected book and actual book are different")
	return nil
}

// BeforeSuite creates a new http library client.
func BeforeSuite() {

	// set the port correctly
	cfg := &client.TransportConfig{
		Host: "localhost:8080",
		BasePath: client.DefaultBasePath,
		Schemes: client.DefaultSchemes,
	}

	HttpClient = client.NewHTTPClientWithConfig(nil, cfg)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(BeforeSuite)
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	ctx.AfterScenario(func(scenario *godog.Scenario, err error){

		// delete the book after the test has run
		params := operations.NewDeleteBookTitleParams()

		params.Title = Book.Title

		if _, err = HttpClient.Operations.DeleteBookTitle(params); err != nil {
			// something has gone wrong, either the test borked,
			// or the cleanup failed.
			panic(err)
		}

	})

	ctx.Step(`^a book is added with title "([^"]*)" and author "([^"]*)"$`, aBookIsAddedWithTitleAndAuthor)
	ctx.Step(`^a library with no books$`, aLibraryWithNoBooks)
	ctx.Step(`^a "([^"]*)" librarian with username "([^"]*)"$`, aLibrarianWithUsername)
	ctx.Step(`^the book can be retrieved$`, theBookCanBeRetrieved)
	ctx.Step(`^a library with book titled "([^"]*)" and author "([^"]*)"$`, aBookIsAddedWithTitleAndAuthor)
	ctx.Step(`^the book is updated with new title "([^"]*)"$`, theBookIsUpdatedWithNewTitle)
	ctx.Step(`^the updated book can be retrieved$`, theBookCanBeRetrieved)
	ctx.Step(`^a book is added with title "([^"]*)", author "([^"]*)" and copies (\d+)$`, aBookIsAddedWithTitleAuthorAndCopies)
	ctx.Step(`^the book has (\d+) copies$`, theBookHasCopies)
	ctx.Step(`^a library with book "([^"]*)", author "([^"]*)" and copies (\d+)$`, aBookIsAddedWithTitleAuthorAndCopies)
	ctx.Step(`^(\d+) more copies of the book are added$`, moreCopiesOfTheBookAreAdded)
}