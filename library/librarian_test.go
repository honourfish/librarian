package library

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"

	"librarian/library/data"
)

// TestAddBook tests Librarian.AddBook in the following ways:
//
// 1. a senior librarian attempts to add a new book successfully.
// 2. a senior librarian attempts to add more stock to an existing book successfully.
// 3. a senior librarian attempts to add a new book, but the persister fails looking for existing books.
// 4. a senior librarian attempts to add a new book, but the persister fails adding the book.
// 5. a senior librarian attempts to add more stock to an existing book, but the persister fails updating the book.
// 6. a non senior librarian attempts to add a new book, but fails with PermissionDeniedError
func TestAddBook(t *testing.T){

	tables := []struct {
		subtest string
		librarian Librarian
		book data.Book

		// Persister stuff
		retrieve_book data.Book

		create_err error
		retrieve_err error
		update_err error
	}{
		{}, // data goes here
	}

	for _, table := range tables {
		t.Run(table.subtest, func(t *testing.T) {

			// assume we have passed a mock persister to librarian
			//table.librarian.Persister.On("Create", "books" table.book).Return(nil)
			
			// mock persister retrieve, to return table.retrieve_book in as the result arg
			//table.librarian.Persister.On("Retrieve", "books", mock.AnythingOfType("interface{}"), mock.AnythingOfType("interface{}")).Return(table.retrieve_err).Run(func(args Arguments){args["result"] = &table.retrieve_book})
			assert.Nil(nil, "")
		})
	}


}