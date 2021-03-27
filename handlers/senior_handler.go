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
	filter := bson.D{{"username", params.Username}}

	var librarian library.Librarian
	if err := slh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewPostLibrarianUsernameBookCreated()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = slh.Persister

	if err := librarian.AddBooks(params.Book.Title, params.Book.Author, int(params.Book.Copies)); err != nil {
		log.Println(err)
		return operations.NewPostLibrarianUsernameBookCreated()
	}

	return operations.NewPostLibrarianUsernameBookCreated()
}