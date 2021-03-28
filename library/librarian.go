package library

import (
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"librarian/library/data"
	"librarian/library/interfaces"
	"librarian/library/errors"
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

	// mutex is the RW lock
	mutex sync.RWMutex
}

// book gets the given book (if any) from persistent storage.
func (l *Librarian) book(title string, author string) (*data.Book, error) {
	var book data.Book

	filter := bson.M{"title":title, "author": author}

	// attempt to get the book from persistent storage
	err := l.Persister.Retrieve("books", filter, &book)

	log.Printf("Get: %d", book.Copies)
	return &book, err
}

// updateBook updates the old_book with the new_book, based on the old_books
//   title and author.
func (l *Librarian) updateBook(old_book *data.Book, new_book *data.Book) (err error) {

	log.Printf("updateBook: %d", new_book.Copies)
	filter := bson.M{"title": old_book.Title, "author": old_book.Author}

	if err = l.Persister.Update("books", filter, new_book); err != nil {
		return
	}

	return
}

// addBook adds a given book to persistent storage
func (l *Librarian) addBook(book *data.Book) (err error) {
	log.Printf("add copies: %d", book.Copies)
	if err = l.Persister.Create("books", book); err != nil {
		return err
	}
	
	return err
}

// removeBook removes a book found with a filter on the title and author.
func (l *Librarian) removeBook(title string, author string) (err error) {

	filter := bson.M{"title": title, "author": author}

	if err = l.Persister.Delete("books", filter); err != nil {
		return
	}

	return
}

// user gets the given user (if any) from persistent storage.
func (l *Librarian) user(username string) (*data.User, error) {
	var user data.User

	filter := bson.M{"username": username}

	// attempt to get the book from persistent storage
	err := l.Persister.Retrieve("users", filter, &user)
	return &user, err
}

// updateUser updates the old_user with the new_user, based on the old_users
//   username.
func (l *Librarian) updateUser(old_user *data.User, new_user *data.User) (err error) {

	filter := bson.M{"username": old_user.Username}

	if err = l.Persister.Update("users", filter, new_user); err != nil {
		return
	}

	return
}

// addUser adds a given user to persistent storage
func (l *Librarian) addUser(user *data.User) (err error) {

	if err = l.Persister.Create("users", user); err != nil {
		return err
	}
	
	return err
}

// removeUser removes a user found with a filter on the usename.
func (l *Librarian) removeUser(username string) (err error) {

	filter := bson.M{"username": username}

	if err = l.Persister.Delete("users", filter); err != nil {
		return
	}

	return
}

// inStock checks if a book is currently in stock.
// Not atomic.
func (l *Librarian) inStock(title string, author string) (bool, error) {

	book, err := l.book(title, author)
	if err != nil {
		return false, err
	}

	if (book.Copies - len(book.Users)) > 0 {
		return true, err
	}

	return false, err
}

// InStock checks if a book is currently in stock.
func (l *Librarian) InStock(title string, author string) (bool, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.inStock(title, author)
}

// Stock gets the current stock information about a book
func (l *Librarian) Stock(title string, author string) (copies int, checked_out int, err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

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
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// first check the book is in stock
	isAvailable, err := l.inStock(title, author)
	if err != nil {
		return err
	}

	if isAvailable == false {
		return &errors.NotEnoughCopiesError{}
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
	user.CheckedOutBooks = append(user.CheckedOutBooks, checked_out_book)
	book.Users = append(book.Users, user.ID)

	if err = l.updateUser(user, user); err != nil {
		 return err
	}
	if err = l.updateBook(book, book); err != nil {
		return err
	}

	return err
}

// CheckIn checks a book back into the library.
func (l *Librarian) CheckIn(title string, author string, username string) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// get user
	var user *data.User
	if user, err = l.user(username); err != nil {
		return
	}

	// get the book
	var book *data.Book
	if book, err = l.book(title, author); err != nil {
		return
	}

	// check the book is checked out by the user
	if !book.ContainsUser(user) {
		return &errors.NotCheckedOutError{}
	}

	// update the book and the user to reflect that the user has checked in the book
	if err = user.RemoveCheckedOutBook(book); err != nil {
		return
	}
	if err = book.RemoveUser(user); err != nil {
		return
	}

	if err = l.updateUser(user, user); err != nil {
		 return
	}
	if err = l.updateBook(book, book); err != nil {
		return
	}

	return
}

// AddBooks adds new stock of a book to the library.
func (l *Librarian) AddBooks(title string, author string, stock int) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// first check we are a senior librarian and have the
	// authorisation needed.
	if l.Role != Senior {
		err = &errors.PermissionDeniedError{Role: l.Role}
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
		return
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
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// first check we are a senior librarian and have the
	// authorisation needed.
	if l.Role != Senior {
		err = &errors.PermissionDeniedError{Role: l.Role}
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
		err = &errors.NotEnoughCopiesError{}
		return
	}

	// if the stock is 0 then remove it
	if (book.Copies - stock) == 0 {
		if err = l.removeBook(title, author); err != nil {
			return
		}
		return
	}

	// otherwise just reduce the number of copies accordingly
	book.Copies -= stock
	if err = l.updateBook(book, book); err != nil {
		return
	}
	
	return
}

// AddUser adds a new user to the registered users.
func (l *Librarian) AddUser(name string, username string) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// attempt to get the user to see if it already exists
	var user *data.User
	user, err = l.user(username)

	// user already exists
	if err == nil {
		return &errors.UserAlreadyExistsError{}
	}

	// retrieving user error
	if err != mongo.ErrNoDocuments {
		return
	}

	user.Name = name
	user.Username = username

	if err = l.addUser(user); err != nil {
		return
	}

	return
}

// RemoveUser removes an existing user from the registered users.
func (l *Librarian) RemoveUser(username string) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	
	err = l.removeUser(username)

	if err == mongo.ErrNoDocuments {
		return &errors.UserDoesNotExistError{}
	}

	// other error unreleated to user not existing
	if err != nil {
		return
	}

	return
}

// User gets the given user (if any) from persistent storage.
func (l *Librarian) User(username string) (user *data.User, err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	user, err = l.user(username)

	// can't find the user
	if err == mongo.ErrNoDocuments {
		return nil, &errors.UserDoesNotExistError{}
	}

	return
}