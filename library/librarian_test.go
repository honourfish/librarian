package library

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/jinzhu/copier"

	"librarian/library/data"
	"librarian/library/mocks"
)

var (
	MockErr error = fmt.Errorf("mock error")
)

// TestAddBooks tests Librarian.AddBook in the following ways:
//
// 1. a senior librarian attempts to add a new book successfully.
// 2. a senior librarian attempts to add more stock to an existing book successfully.
// 3. a senior librarian attempts to add a new book, but the persister fails looking for existing books.
// 4. a senior librarian attempts to add a new book, but the persister fails adding the book.
// 5. a senior librarian attempts to add more stock to an existing book, but the persister fails updating the book.
// 6. a non senior librarian attempts to add a new book, but fails with PermissionDeniedError
func TestAddBooks(t *testing.T){

	tables := []struct {
		subtest string
		librarian Librarian
		title string
		author string
		stock int

		// Persister stuff
		retrieve_book data.Book

		create_err error
		retrieve_err error
		update_err error

		expected error
	}{
		{
			"senior create successful",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{},
			nil, // create error
			mongo.ErrNoDocuments, // retrieve error
			nil, // update error

			// expected error
			nil,
		},
		{
			"senior update successful",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 5,
			},
			nil, // create error
			nil, // retrieve error
			nil, // update error

			// expected error
			nil,
		},
		{
			"senior retrieve error",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{},
			nil, // create error
			MockErr, // retrieve error
			nil, // update error

			// expected error
			MockErr,
		},
		{
			"senior create error",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{},
			MockErr, // create error
			mongo.ErrNoDocuments, // retrieve error
			nil, // update error

			// expected error
			MockErr,
		},
		{
			"senior update error",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{},
			nil, // create error
			nil, // retrieve error
			MockErr, // update error

			// expected error
			MockErr,
		},
		{
			"librarian denied",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{},
			nil, // create error
			nil, // retrieve error
			nil, // update error

			// expected error
			&PermissionDeniedError{Role: "Librarian"},
		},
	}

	for _, table := range tables {
		t.Run(table.subtest, func(t *testing.T) {

			// assume we have passed a mock persister to librarian
			table.librarian.Persister.(*mocks.MockPersister).On("Create", "books", mock.Anything).Return(table.create_err)
	
			table.librarian.Persister.(*mocks.MockPersister).On("Update", "books", mock.Anything, mock.Anything).Return(table.update_err)

			// mock persister retrieve, to return table.retrieve_book in as the result arg
			table.librarian.Persister.(*mocks.MockPersister).On("Retrieve", "books", mock.Anything, mock.Anything).Return(table.retrieve_err).Run(func(args mock.Arguments) {
				arg := args.Get(2).(*data.Book)
				copier.Copy(arg, &table.retrieve_book)
			})

			actual := table.librarian.AddBooks(table.title, table.author, table.stock)

			assert.Equal(t, actual, table.expected, "error returned not the same as eexpppected")
		})
	}
}

// TestRemoveBooks tests Librarian.RemoveBooks in the following ways:
//
// 1. a senior librarian attempts to remove book copies successfully, leaving more than one copy.
// 2. a senior librarian attempts to remove book copies successfully, leaving no copies.
// 3. a senior librarian attempts to remove a book, but the persister fails looking for existing books.
// 4. a senior librarian attempts to remove all book copies, but the persister fails deleting the book.
// 5. a senior librarian attempts to remove some book copies, leaving more than one copy, but the persister fails updating the book.
// 6. a senior librarian attempts to remove more copies of the book than exist, fails with NotEnoughCopiesError.
// 7. a non senior librarian attempts to remove a new book, but fails with PermissionDeniedError.
func TestRemoveBooks(t *testing.T){

	tables := []struct {
		subtest string
		librarian Librarian
		title string
		author string
		stock int

		// Persister stuff
		retrieve_book data.Book

		delete_err error
		retrieve_err error
		update_err error

		expected error
	}{
		{
			"senior remove some successful",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
			},
			nil, // delete error
			nil, // retrieve error
			nil, // update error

			// expected error
			nil,
		},
		{
			"senior remove all successful",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 5,
			},
			nil, // delete error
			nil, // retrieve error
			nil, // update error

			// expected error
			nil,
		},
		{
			"senior retrieve error",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{},
			nil, // delete error
			MockErr, // retrieve error
			nil, // update error

			// expected error
			MockErr,
		},
		{
			"senior remove error",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 5,
			},
			MockErr, // delete error
			nil, // retrieve error
			nil, // update error

			// expected error
			MockErr,
		},
		{
			"senior update error",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
			},
			nil, // delete error
			nil, // retrieve error
			MockErr, // update error

			// expected error
			MockErr,
		},
		{
			"senior not enough copies error",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 3,
			},
			nil, // delete error
			nil, // retrieve error
			nil, // update error

			// expected error
			&NotEnoughCopiesError{},
		},
		{
			"librarian denied",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			5, // stock

			data.Book{},
			nil, // delete error
			nil, // retrieve error
			nil, // update error

			// expected error
			&PermissionDeniedError{Role: "Librarian"},
		},
	}

	for _, table := range tables {
		t.Run(table.subtest, func(t *testing.T) {

			// assume we have passed a mock persister to librarian
			table.librarian.Persister.(*mocks.MockPersister).On("Delete", "books", mock.Anything).Return(table.delete_err)
	
			table.librarian.Persister.(*mocks.MockPersister).On("Update", "books", mock.Anything, mock.Anything).Return(table.update_err)

			// mock persister retrieve, to return table.retrieve_book in as the result arg
			table.librarian.Persister.(*mocks.MockPersister).On("Retrieve", "books", mock.Anything, mock.Anything).Return(table.retrieve_err).Run(func(args mock.Arguments) {
				arg := args.Get(2).(*data.Book)
				copier.Copy(arg, &table.retrieve_book)
			})

			actual := table.librarian.RemoveBooks(table.title, table.author, table.stock)

			assert.Equal(t, actual, table.expected, "error returned not the same as eexpppected")
		})
	}
}