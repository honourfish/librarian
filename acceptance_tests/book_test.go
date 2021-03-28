package main

import (
	"fmt"
	"github.com/cucumber/godog"

	//"github.com/stretchr/testify/assert"

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
var BookStock *models.BookStock
var t asserter

func aLibraryWithNoBooks() error {
	// nothing in the library yet
	return nil
}

func theBookCanBeRetrieved() error {
	// use default constuctor to set the default request timeout
	params := operations.NewGetLibrarianUsernameBookTitleAuthorParams()

	params.Username = Librarian.Username
	params.Title = Book.Title
	params.Author = Book.Author

	response, err := HttpClient.Operations.GetLibrarianUsernameBookTitleAuthor(params)
	if err != nil {
		return err
	}

	var bookStock models.BookStock = *response.Payload
	BookStock = &bookStock
	return t.err
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

// containsTag checks if a scenario contains a given tag, e.g. '@book'.
func containsTag(tag string, scenario *godog.Scenario) bool {
    for _, scenario_tag := range scenario.Tags {
        if scenario_tag.Name == tag {
            return true
        }
    }
    return false
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	ctx.AfterScenario(func(scenario *godog.Scenario, err error){

		if containsTag("@book", scenario) {
			// delete the book after the test has run
			params := operations.NewDeleteBookTitleParams()

			params.Title = Book.Title

			if _, err = HttpClient.Operations.DeleteBookTitle(params); err != nil {
				// something has gone wrong, either the test borked,
				// or the cleanup failed.
				panic(err)
			}
		}

		if containsTag("@user", scenario) {
			// delete the user after the test has run
			params := operations.NewDeleteLibrarianUsernameUserUserParams()

			params.Username = Librarian.Username
			params.User = User.Username

			if _, err = HttpClient.Operations.DeleteLibrarianUsernameUserUser(params); err != nil {
				// something has gone wrong, either the test borked,
				// or the cleanup failed.
				panic(err)
			}
		}

	})
	ctx.Step(`^a library with no books$`, aLibraryWithNoBooks)
	ctx.Step(`^a "([^"]*)" librarian with username "([^"]*)"$`, aLibrarianWithUsername)
	ctx.Step(`^the book can be retrieved$`, theBookCanBeRetrieved)
	ctx.Step(`^the updated book can be retrieved$`, theBookCanBeRetrieved)
	ctx.Step(`^a book is added with title "([^"]*)", author "([^"]*)" and copies (\d+)$`, aBookIsAddedWithTitleAuthorAndCopies)
	ctx.Step(`^the book has (\d+) copies$`, theBookHasCopies)
	ctx.Step(`^a library with book "([^"]*)", author "([^"]*)" and copies (\d+)$`, aLibraryWithBookAuthorAndCopies)
	ctx.Step(`^(\d+) more copies of the book are added$`, moreCopiesOfTheBookAreAdded)
	ctx.Step(`^(\d+) copies of the book are removed$`, copiesOfTheBookAreRemoved)
	ctx.Step(`^a library with no users$`, aLibraryWithNoUsers)
	ctx.Step(`^a user is added with username "([^"]*)" and name "([^"]*)"$`, aUserIsAddedWithUsernameAndName)
	ctx.Step(`^the user can be retrieved$`, theUserCanBeRetrieved)
	ctx.Step(`^a library with user username "([^"]*)" and name "([^"]*)"$`, aLibraryWithUserUsernameAndName)
	ctx.Step(`^the checked out copies of the book is (\d+)$`, theCheckedOutCopiesOfTheBookIs)
	ctx.Step(`^the user checks out the book$`, theUserChecksOutTheBook)
	ctx.Step(`^the user checks in the book$`, theUserChecksInTheBook)
}