package handlers

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/go-openapi/runtime/middleware"

	"librarian/swagger/restapi/operations"
	//"librarian/swagger/models"

	"librarian/library"
	//"librarian/library/data"
	"librarian/library/interfaces"
)

type SeniorLibrarianHandler struct {
	Persister interfaces.Persister
}


// handlePostBook is the handler for post request for the endpoint /librarian/{username}/book
func (slh *SeniorLibrarianHandler) HandlePostBook(params operations.PostLibrarianUsernameBookParams) middleware.Responder {

	// get the librarian
	filter := bson.M{"username": params.Username}

	var librarian library.Librarian
	if err := slh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewPostLibrarianUsernameBookNotFound()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = slh.Persister

	if err := librarian.AddBooks(params.Book.Title, params.Book.Author, int(params.Book.Copies)); err != nil {
		log.Println(err)
		return operations.NewPostLibrarianUsernameBookNotFound()
	}

	return operations.NewPostLibrarianUsernameBookCreated()
}

// handleDeleteBook is the handler for delete request for the endpoint /librarian/{username}/book/{title}/{author}/{copies}
func (slh *SeniorLibrarianHandler) HandleDeleteBook(params operations.DeleteLibrarianUsernameBookTitleAuthorCopiesParams) middleware.Responder {

	// get the librarian
	filter := bson.M{"username": params.Username}

	var librarian library.Librarian
	if err := slh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewDeleteLibrarianUsernameBookTitleAuthorCopiesNotFound()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = slh.Persister

	if err := librarian.RemoveBooks(params.Title, params.Author, int(params.Copies)); err != nil {
		log.Println(err)
		return operations.NewDeleteLibrarianUsernameBookTitleAuthorCopiesNotFound()
	}

	return operations.NewDeleteLibrarianUsernameBookTitleAuthorCopiesOK()
}