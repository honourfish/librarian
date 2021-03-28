package library

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jinzhu/copier"

	"librarian/library/data"
	"librarian/library/mocks"
	"librarian/library/errors"
)

var (
	MockErr error = fmt.Errorf("mock error")
)

// NewObjectID is constructor for the custom ObjectID type
func NewObjectID(id byte) primitive.ObjectID {
	var object_id primitive.ObjectID
	object_id = [12]byte{id}
	return object_id
}

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
			&errors.PermissionDeniedError{Role: "Librarian"},
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
			&errors.NotEnoughCopiesError{},
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
			&errors.PermissionDeniedError{Role: "Librarian"},
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

// TestCheckout tests Librarian.Checkout in the following ways:
//
// 1. a senior librarian attempts to checkout a book to a registered user, successfully.
// 2. a normal librarian attempts to checkout a book to a registered user, successfully.
// 3. a normal librarian attempts to checkout a book to a non registered user, unsuccessfully.
// 4. a normal librarian attempts to checkout a non existing book to a registered user, unsuccessfully.
// 5. a normal librarian attempts to checkout a book that has all its copies checkout out already to a registered user, unsuccessfully.
// 6. a normal librarian attempts to checkout a book to a registered user, but the user update fails.
// 7. a normal librarian attempts to checkout a book to a registered user, but the book update fails.
func TestCheckOut(t *testing.T){

	tables := []struct {
		subtest string
		librarian Librarian
		title string
		author string
		username string

		// Persister stuff
		retrieve_book data.Book
		retrieve_user data.User

		retrieve_book_err error
		retrieve_user_err error
		update_book_err error
		update_user_err error

		expected error
	}{
		{
			"senior checkout successful",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
			}, // retrieve book

			data.User{
				Username: "Phil",
			}, // retrieve user

			nil, // retrieve book error
			nil, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			nil,
		},
		{
			"normal checkout successful",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
			}, // retrieve book

			data.User{
				Username: "Phil",
			}, // retrieve user

			nil, // retrieve book error
			nil, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			nil,
		},
		{
			"normal non registered user",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
			}, // retrieve book

			data.User{}, // retrieve user

			nil, // retrieve book error
			MockErr, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			MockErr,
		},
		{
			"normal non existent book",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{}, // retrieve book

			data.User{}, // retrieve user

			MockErr, // retrieve book error
			nil, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			MockErr,
		},
		{
			"normal all checked out",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 1, // only 1 copy of the book available
				Users: []primitive.ObjectID{NewObjectID(1)}, // 1 book checked out
			}, // retrieve book

			data.User{
				Username: "Phil",
			}, // retrieve user

			nil, // retrieve book error
			nil, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			&errors.NotEnoughCopiesError{},
		},
		{
			"normal user update failed",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 1, 
			}, // retrieve book

			data.User{
				Username: "Phil",
			}, // retrieve user

			nil, // retrieve book error
			MockErr, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			MockErr,
		},
		{
			"normal book update failed",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 1, 
			}, // retrieve book

			data.User{
				Username: "Phil",
			}, // retrieve user

			nil, // retrieve book error
			nil, // retrieve user error
			MockErr, // update book error
			nil, // update user error

			// expected error
			MockErr,
		},
	}

	for _, table := range tables {
		t.Run(table.subtest, func(t *testing.T) {

			// mock stuff
			table.librarian.Persister.(*mocks.MockPersister).On("Update", "books", mock.Anything, mock.Anything).Return(table.update_book_err)
			table.librarian.Persister.(*mocks.MockPersister).On("Update", "users", mock.Anything, mock.Anything).Return(table.update_user_err)

			// mock persister retrieve, to return table.retrieve_book in as the result arg
			table.librarian.Persister.(*mocks.MockPersister).On("Retrieve", "books", mock.Anything, mock.Anything).Return(table.retrieve_book_err).Run(func(args mock.Arguments) {
				arg := args.Get(2).(*data.Book)
				copier.Copy(arg, &table.retrieve_book)
			})
			table.librarian.Persister.(*mocks.MockPersister).On("Retrieve", "users", mock.Anything, mock.Anything).Return(table.retrieve_user_err).Run(func(args mock.Arguments) {
				arg := args.Get(2).(*data.User)
				copier.Copy(arg, &table.retrieve_user)
			})

			actual := table.librarian.CheckOut(table.title, table.author, table.username)

			assert.Equal(t, table.expected, actual, "error returned not the same as expected")
		})
	}
}

// TestCheckIn tests Librarian.CheckIn in the following ways:
//
// 1. a senior librarian attempts to checkin a book from a registered user, successfully.
// 2. a normal librarian attempts to checkin a book from a registered user, successfully.
// 3. a normal librarian attempts to checkin a book from a non registered user, unsuccessfully.
// 4. a normal librarian attempts to checkin a non existing book from a registered user, unsuccessfully.
// 5. a normal librarian attempts to checkin a book that the registered user hasn't checked out, unsuccessfully.
// 6. a normal librarian attempts to checkin a book from a registered user, but the user update fails.
// 7. a normal librarian attempts to checkin a book from a registered user, but the book update fails.
func TestCheckIn(t *testing.T){

	tables := []struct {
		subtest string
		librarian Librarian
		title string
		author string
		username string

		// Persister stuff
		retrieve_book data.Book
		retrieve_user data.User

		retrieve_book_err error
		retrieve_user_err error
		update_book_err error
		update_user_err error

		expected error
	}{
		{
			"senior checkin successful",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
				Users: []primitive.ObjectID{
					NewObjectID(2),
				},
			}, // retrieve book

			data.User{
				ID: NewObjectID(2),
				Username: "Phil",
				CheckedOutBooks: []*data.CheckedOutBook{
					&data.CheckedOutBook{
						BookRef: NewObjectID(1),
					},
				},
			}, // retrieve user

			nil, // retrieve book error
			nil, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			nil,
		},
		{
			"normal checkin successful",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
				Users: []primitive.ObjectID{
					NewObjectID(2),
				},
			}, // retrieve book

			data.User{
				ID: NewObjectID(2),
				Username: "Phil",
				CheckedOutBooks: []*data.CheckedOutBook{
					&data.CheckedOutBook{
						BookRef: NewObjectID(1),
					},
				},
			}, // retrieve user

			nil, // retrieve book error
			nil, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			nil,
		},
		{
			"normal non registered user",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
				Users: []primitive.ObjectID{
					NewObjectID(2),
				},
			}, // retrieve book

			data.User{}, // retrieve user

			nil, // retrieve book error
			MockErr, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			MockErr,
		},
		{
			"normal non existent book",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{}, // retrieve book

			data.User{
				ID: NewObjectID(2),
				Username: "Phil",
				CheckedOutBooks: []*data.CheckedOutBook{
					&data.CheckedOutBook{
						BookRef: NewObjectID(1),
					},
				},
			}, // retrieve user

			MockErr, // retrieve book error
			nil, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			MockErr,
		},
		{
			"normal checkin non checked out book",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
			}, // retrieve book

			data.User{
				ID: NewObjectID(2),
				Username: "Phil",
			}, // retrieve user

			nil, // retrieve book error
			nil, // retrieve user error
			nil, // update book error
			nil, // update user error

			// expected error
			&errors.NotCheckedOutError{},
		},
		{
			"normal update user failed",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
				Users: []primitive.ObjectID{
					NewObjectID(2),
				},
			}, // retrieve book

			data.User{
				ID: NewObjectID(2),
				Username: "Phil",
				CheckedOutBooks: []*data.CheckedOutBook{
					&data.CheckedOutBook{
						BookRef: NewObjectID(1),
					},
				},
			}, // retrieve user

			nil, // retrieve book error
			nil, // retrieve user error
			nil, // update book error
			MockErr, // update user error

			// expected error
			MockErr,
		},
		{
			"normal update book failed",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author
			"Phil", // username

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
				Users: []primitive.ObjectID{
					NewObjectID(2),
				},
			}, // retrieve book

			data.User{
				ID: NewObjectID(2),
				Username: "Phil",
				CheckedOutBooks: []*data.CheckedOutBook{
					&data.CheckedOutBook{
						BookRef: NewObjectID(1),
					},
				},
			}, // retrieve user

			nil, // retrieve book error
			nil, // retrieve user error
			MockErr, // update book error
			nil, // update user error

			// expected error
			MockErr,
		},
	}

	for _, table := range tables {
		t.Run(table.subtest, func(t *testing.T) {

			// mock stuff
			table.librarian.Persister.(*mocks.MockPersister).On("Update", "books", mock.Anything, mock.Anything).Return(table.update_book_err)
			table.librarian.Persister.(*mocks.MockPersister).On("Update", "users", mock.Anything, mock.Anything).Return(table.update_user_err)

			// mock persister retrieve, to return table.retrieve_book in as the result arg
			table.librarian.Persister.(*mocks.MockPersister).On("Retrieve", "books", mock.Anything, mock.Anything).Return(table.retrieve_book_err).Run(func(args mock.Arguments) {
				arg := args.Get(2).(*data.Book)
				copier.Copy(arg, &table.retrieve_book)
			})
			table.librarian.Persister.(*mocks.MockPersister).On("Retrieve", "users", mock.Anything, mock.Anything).Return(table.retrieve_user_err).Run(func(args mock.Arguments) {
				arg := args.Get(2).(*data.User)
				copier.Copy(arg, &table.retrieve_user)
			})

			actual := table.librarian.CheckIn(table.title, table.author, table.username)

			assert.Equal(t, table.expected, actual, "error returned not the same as expected")
		})
	}
}

// TestInStock tests Librarian.InStock in the following ways:
//
// 1. a senior librarian attempts to check stock of a book with multiple copies, but none checked out, successfully and returns true.
// 2. a normal librarian attempts to check stock of a book with multiple copies, but none checked out, successfully and returns true.
// 3. a normal librarian attempts to check stock of a book with no copies and none checked out, successfully and returns false.
// 4. a normal librarian attempts to check stock of a book with multiple copies and all checked out, successfully and returns false.
// 5. a normal librarian attempts to check stock of a book but retrieving the book fails and returns false.
func TestInStock(t *testing.T){

	tables := []struct {
		subtest string
		librarian Librarian
		title string
		author string

		// Persister stuff
		retrieve_book data.Book

		retrieve_book_err error

		expected bool
		expected_err error
	}{
		{
			"senior in stock successful",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
				Users: []primitive.ObjectID{
					NewObjectID(2),
				},
			}, // retrieve book

			nil, // retrieve book error

			// expected
			true,
			nil,
		},
		{
			"normal in stock successful",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
				Users: []primitive.ObjectID{
					NewObjectID(2),
				},
			}, // retrieve book

			nil, // retrieve book error

			// expected
			true,
			nil,
		},
		{
			"normal no copies",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 0,
			}, // retrieve book

			nil, // retrieve book error

			// expected
			false,
			nil,
		},
		{
			"normal all checked out",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 2,
				Users: []primitive.ObjectID{
					NewObjectID(2),
					NewObjectID(3),
				},
			}, // retrieve book

			nil, // retrieve book error

			// expected
			false,
			nil,
		},
		{
			"normal retrieve book failed",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 4,
				Users: []primitive.ObjectID{
					NewObjectID(2),
					NewObjectID(3),
				},
			}, // retrieve book

			MockErr, // retrieve book error

			// expected
			false,
			MockErr,
		},
	}

	for _, table := range tables {
		t.Run(table.subtest, func(t *testing.T) {

			// mock persister retrieve, to return table.retrieve_book in as the result arg
			table.librarian.Persister.(*mocks.MockPersister).On("Retrieve", "books", mock.Anything, mock.Anything).Return(table.retrieve_book_err).Run(func(args mock.Arguments) {
				arg := args.Get(2).(*data.Book)
				copier.Copy(arg, &table.retrieve_book)
			})

			actual, actual_err := table.librarian.InStock(table.title, table.author)

			assert.Equal(t, table.expected, actual, "bool returned not the same as expected")
			assert.Equal(t, table.expected_err, actual_err, "error returned not the same as expected")
		})
	}
}

// TestStock tests Librarian.Stock in the following ways:
//
// 1. a senior librarian attempts to take stock of a book with multiple copies, but none checked out, successfully.
// 2. a normal librarian attempts to take stock of a book with multiple copies, but none checked out, successfully.
// 3. a normal librarian attempts to take stock of a book with multiple copies, but some checked out, successfully.
// 4. a normal librarian attempts to take stock of a book but retrieving the book fails and returns false.
func TestStock(t *testing.T){

	tables := []struct {
		subtest string
		librarian Librarian
		title string
		author string

		// Persister stuff
		retrieve_book data.Book

		retrieve_book_err error

		expected_copies int
		expected_checked_out int
		expected_err error
	}{
		{
			"senior take stock successful",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 10,
			}, // retrieve book

			nil, // retrieve book error

			// expected
			10,
			0,
			nil,
		},
		{
			"normal take stock successful",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 7,
			}, // retrieve book

			nil, // retrieve book error

			// expected
			7,
			0,
			nil,
		},
		{
			"normal take stock some checked out",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Harry Potter", // title
			"J.K. Rowling", // author

			data.Book{
				ID: NewObjectID(1),
				Title: "Harry Potter",
				Author: "J.K. Rowling",
				Copies: 5,
				Users: []primitive.ObjectID{
					NewObjectID(2),
					NewObjectID(3),
				},
			}, // retrieve book

			nil, // retrieve book error

			// expected
			5,
			2,
			nil,
		},
	}

	for _, table := range tables {
		t.Run(table.subtest, func(t *testing.T) {

			// mock persister retrieve, to return table.retrieve_book in as the result arg
			table.librarian.Persister.(*mocks.MockPersister).On("Retrieve", "books", mock.Anything, mock.Anything).Return(table.retrieve_book_err).Run(func(args mock.Arguments) {
				arg := args.Get(2).(*data.Book)
				copier.Copy(arg, &table.retrieve_book)
			})

			actual_copies, actual_checked_out, actual_err := table.librarian.Stock(table.title, table.author)

			assert.Equal(t, table.expected_copies, actual_copies, "copies returned not the same as expected")
			assert.Equal(t, table.expected_checked_out, actual_checked_out, "checked out returned not the same as expected")
			assert.Equal(t, table.expected_err, actual_err, "error returned not the same as expected")
		})
	}
}

// TestAddUser tests Librarian.AddUser in the following ways:
//
// 1. a senior librarian attempts to add a new user successfully.
// 2. a normal librarian attempts to add a new user successfully.
// 3. a normal librarian attempts to add an existing user unsuccessfully.
// 4. a normal librarian attempts to add a new user, but retrieve user fails.
// 5. a normal librarian attempts to add a new user, but add user fails.
func TestAddUser(t *testing.T){

	tables := []struct {
		subtest string
		librarian Librarian
		username string

		// Persister stuff
		retrieve_user data.User

		create_err error
		retrieve_err error

		expected error
	}{
		{
			"senior create successful",
			Librarian{
				Name: "Janice",
				Role: "Senior",
				Persister: &mocks.MockPersister{},
			},
			"Phil", // username

			data.User{}, // retrieve user

			nil, // create error
			mongo.ErrNoDocuments, // retrieve error

			// expected error
			nil,
		},
		{
			"normal create successful",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Phil", // username

			data.User{}, // retrieve user

			nil, // create error
			mongo.ErrNoDocuments, // retrieve error

			// expected error
			nil,
		},
		{
			"normal existing user",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Phil", // username

			data.User{
				Username: "Phil",
			}, // retrieve user

			nil, // create error
			nil, // retrieve error

			// expected error
			&errors.UserAlreadyExistsError{},
		},
		{
			"normal retrieve failed",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Phil", // username

			data.User{}, // retrieve user

			nil, // create error
			MockErr, // retrieve error

			// expected error
			MockErr,
		},
		{
			"normal create failed",
			Librarian{
				Name: "Janice",
				Role: "Librarian",
				Persister: &mocks.MockPersister{},
			},
			"Phil", // username

			data.User{}, // retrieve user

			MockErr, // create error
			mongo.ErrNoDocuments, // retrieve error

			// expected error
			MockErr,
		},
	}

	for _, table := range tables {
		t.Run(table.subtest, func(t *testing.T) {

			// assume we have passed a mock persister to librarian
			table.librarian.Persister.(*mocks.MockPersister).On("Create", "users", mock.Anything).Return(table.create_err)
	
			// mock persister retrieve, to return table.retrieve_book in as the result arg
			table.librarian.Persister.(*mocks.MockPersister).On("Retrieve", "users", mock.Anything, mock.Anything).Return(table.retrieve_err).Run(func(args mock.Arguments) {
				arg := args.Get(2).(*data.User)
				copier.Copy(arg, &table.retrieve_user)
			})

			actual := table.librarian.AddUser("Philly", table.username)

			assert.Equal(t, actual, table.expected, "error returned not the same as expected")
		})
	}
}