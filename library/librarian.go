package library

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"librarian/library/data"
	"librarian/library/interfaces"
)

const (
	Normal string = "Librarian"
	Senior string = "Senior"
)

// Librarian is a library administrator that depending on role can
//   checkin or checkout books.
type Librarian struct {
	// ID is the unique mongodb id.
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Name of the librarian. Non unique
	Name string `bson:"name,omitempty"`

	// Username of the librarian. Unique
	Username string `bson:"username,omitempty"`

	// Role is the librarians role (use constant Normal or Senior).
	Role string `bson:"role,omitempty"`

	// Persister is the persister used to add users
	//   and books to persistant storage.
	Persister interfaces.Persister
}

// book gets the given book (if any) from persistent storage.
func (l *Librarian) book(title string, author string) (*data.Book, error) {
	var book data.Book

	filter := bson.D{{"title", title},{"author", author}}

	// attempt to get the book from persistent storage
	err := l.Persister.Retrieve("books", filter, &book)
	return &book, err
}

// updateBook updates the old_book with the new_book, based on the old_books
//   title and author.
func (l *Librarian) updateBook(old_book *data.Book, new_book *data.Book) (err error) {

	filter := bson.D{{"title", old_book.Title},{"author", old_book.Author}}

	if err = l.Persister.Update("books", filter, new_book); err != nil {
		return
	}

	return
}

// addBook adds a given book to persistent storage
func (l *Librarian) addBook(book *data.Book) (err error) {

	if err = l.Persister.Create("books", book); err != nil {
		return err
	}
	
	return err
}

// removeBook removes a book found with a filter on the title and author.
func (l *Librarian) removeBook(title string, author string) (err error) {

	filter := bson.D{{"title", title},{"author", author}}

	if err = l.Persister.Delete("books", filter); err != nil {
		return
	}

	return
}

// user gets the given user (if any) from persistent storage.
func (l *Librarian) user(username string) (*data.User, error) {
	var user data.User

	filter := bson.D{{"username", username}}

	// attempt to get the book from persistent storage
	err := l.Persister.Retrieve("users", filter, &user)
	return &user, err
}

// inStock checks if a book is currently in stock.
// Not atomic.
func (l *Librarian) inStock(title string, author string) (bool, error) {

	book, err := l.book(title, author)
	if err != nil {
		return false, err
	}

	if book.Copies > 0 {
		return true, err
	}
	return false, err
}

// InStock checks if a book is currently in stock.
// Todo: make public method atomic.
func (l *Librarian) InStock(title string, author string) (bool, error) {
	return l.inStock(title, author)
}

// Stock gets the current stock information about a book
func (l *Librarian) Stock(title string, author string) (copies int, checked_out int, err error) {
	var book *data.Book

	book, err = l.book(title, author)
	if err != nil {
		return
	}

	copies = book.Copies
	checked_out = len(book.Users)
	return
}

// CheckOut checks a book in the library out.
func (l *Librarian) CheckOut(title string, author string, username string) error {

	// first check the book is in stock
	isAvailable, err := l.inStock(title, author)
	if err != nil {
		return err
	}

	if isAvailable == false {
		return &NotEnoughCopiesError{}
	}

	// get user
	var user *data.User
	if user, err = l.user(username); err != nil {
		return err
	}

	// get the book
	var book *data.Book
	if book, err = l.book(title, author); err != nil {
		return err
	}

	// create a checked out book
	checked_out_book := &data.CheckedOutBook{
		BookRef: book.ID,
		CheckOutDate: time.Now(),
		DueDate: time.Now(), // todo one week in future
	}

	// update the book and the user to reflect that the user has checked out the book
	user.CheckedOutBooks = append(user.CheckedOutBooks, *checked_out_book)
	book.Users = append(book.Users, user.ID)

	//if err = l.updateUser(user, user); err != nil {
		// return err
	//}
	if err = l.updateBook(book, book); err != nil {
		return err
	}

	return err
}

// AddBooks adds new stock of a book to the library.
func (l *Librarian) AddBooks(title string, author string, stock int) (err error) {

	// first check we are a senior librarian and have the
	// authorisation needed.
	if l.Role != Senior {
		err = &PermissionDeniedError{Role: l.Role}
		return
	}

	// attempt to get the book to see if it already exists
	var book *data.Book
	book, err = l.book(title, author)

	// the book doesn't exist so create it
	if err == mongo.ErrNoDocuments {
		book.Title = title
		book.Author = author
		book.Copies = stock

		if err = l.addBook(book); err != nil {
			return
		}
	}

	// check there are no other errors getting the book
	if err != nil {
		return
	}

	// book exists so update it
	book.Copies += stock
	if err = l.updateBook(book, book); err != nil {
		return
	}

	return
}

// RemoveBooks removes stock of a book from the library.
func (l *Librarian) RemoveBooks(title string, author string, stock int) (err error) {

	// first check we are a senior librarian and have the
	// authorisation needed.
	if l.Role != Senior {
		err = &PermissionDeniedError{Role: l.Role}
		return
	}

	// attempt to get the book to see if it already exists
	var book *data.Book
	book, err = l.book(title, author)

	// check we have a book to remove
	if err != nil {
		return
	}

	// check we have enough copies to remove
	if (book.Copies - stock) < 0 {
		err = &NotEnoughCopiesError{}
		return
	}

	// if the stock is 0 then remove it
	if (book.Copies - stock) == 0 {
		if err = l.removeBook(title, author); err != nil {
			return
		}
	}

	// otherwise just reduce the number of copies accordingly
	book.Copies -= stock
	if err = l.updateBook(book, book); err != nil {
		return
	}
	
	return
}