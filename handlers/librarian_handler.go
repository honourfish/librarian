package handlers

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/go-openapi/runtime/middleware"

	"librarian/swagger/restapi/operations"
	"librarian/swagger/models"

	"librarian/library"
	//"librarian/library/data"
	"librarian/library/interfaces"
)

type LibrarianHandler struct {
	Persister interfaces.Persister
}


// handlePostUser is the handler for post request for the endpoint /librarian/{username}/user
func (lh *LibrarianHandler) HandlePostUser(params operations.PostLibrarianUsernameUserParams) middleware.Responder {

	// get the librarian
	filter := bson.M{"username": params.Username}

	var librarian library.Librarian
	if err := lh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewPostLibrarianUsernameUserNotFound()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = lh.Persister

	if err := librarian.AddUser(params.User.Name, params.User.Username); err != nil {
		log.Println(err)
		return operations.NewPostLibrarianUsernameUserNotFound()
	}

	return operations.NewPostLibrarianUsernameUserCreated()
}

// handleDeleteUser is the handler for delete request for the endpoint /librarian/{username}/user/{user}
func (lh *LibrarianHandler) HandleDeleteUser(params operations.DeleteLibrarianUsernameUserUserParams) middleware.Responder {

	// get the librarian
	filter := bson.M{"username": params.Username}

	var librarian library.Librarian
	if err := lh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewDeleteLibrarianUsernameUserUserNotFound()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = lh.Persister

	if err := librarian.RemoveUser(params.User); err != nil {
		log.Println(err)
		return operations.NewDeleteLibrarianUsernameUserUserNotFound()
	}

	return operations.NewDeleteLibrarianUsernameUserUserOK()
}

// HandleGetUser is the handler for get request for the endpoint /librarian/{username}/user/{user}
func (lh *LibrarianHandler) HandleGetUser(params operations.GetLibrarianUsernameUserUserParams) middleware.Responder {

	// get the librarian
	filter := bson.M{"username": params.Username}

	var librarian library.Librarian
	if err := lh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewGetLibrarianUsernameUserUserNotFound()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = lh.Persister

	user, err := librarian.User(params.User)
	if err != nil {
		log.Println(err)
		return operations.NewGetLibrarianUsernameUserUserNotFound()
	}

	responseUser := &models.User{
		Name: user.Name,
		Username: user.Username,
	}

	return operations.NewGetLibrarianUsernameUserUserOK().WithPayload(responseUser)
}

// HandleGetBook is the handler for get request for the endpoint /librarian/{username}/book/{title}/{author}
func (lh *LibrarianHandler) HandleGetBook(params operations.GetLibrarianUsernameBookTitleAuthorParams) middleware.Responder {

	// get the librarian
	filter := bson.M{"username": params.Username}

	var librarian library.Librarian
	if err := lh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewGetLibrarianUsernameBookTitleAuthorNotFound()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = lh.Persister

	copies, checked_out, err := librarian.Stock(params.Title, params.Author)
	if err != nil {
		log.Println(err)
		return operations.NewGetLibrarianUsernameBookTitleAuthorNotFound()
	}

	var in_stock bool
	in_stock, err = librarian.InStock(params.Title, params.Author)
	if err != nil {
		log.Println(err)
		return operations.NewGetLibrarianUsernameBookTitleAuthorNotFound()
	}

	responseBookStock := &models.BookStock{
		Copies: int64(copies),
		Checkedout: int64(checked_out),
		Instock: in_stock,
	}

	return operations.NewGetLibrarianUsernameBookTitleAuthorOK().WithPayload(responseBookStock)
}

// HandlePutCheckOut is the handler for put request for the endpoint /librarian/{username}/user/{user}/checkout
func (lh *LibrarianHandler) HandlePutCheckOut(params operations.PutLibrarianUsernameUserUserCheckoutParams) middleware.Responder {

	// get the librarian
	filter := bson.M{"username": params.Username}

	var librarian library.Librarian
	if err := lh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewPutLibrarianUsernameUserUserCheckoutNotFound()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = lh.Persister

	if err := librarian.CheckOut(*params.Book.Title, *params.Book.Author, params.User); err != nil {
		log.Println(err)
		return operations.NewPutLibrarianUsernameUserUserCheckoutNotFound()
	}

	return operations.NewPutLibrarianUsernameUserUserCheckoutOK()
}