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
	filter := bson.D{{"username", params.Username}}

	var librarian library.Librarian
	if err := lh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewPostLibrarianUsernameUserCreated()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = lh.Persister

	if err := librarian.AddUser(params.User.Name, params.User.Username); err != nil {
		log.Println(err)
		return operations.NewPostLibrarianUsernameUserCreated()
	}

	return operations.NewPostLibrarianUsernameUserCreated()
}

// handleDeleteUser is the handler for delete request for the endpoint /librarian/{username}/user/{user}
func (lh *LibrarianHandler) HandleDeleteUser(params operations.DeleteLibrarianUsernameUserUserParams) middleware.Responder {

	// get the librarian
	filter := bson.D{{"username", params.Username}}

	var librarian library.Librarian
	if err := lh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewDeleteLibrarianUsernameUserUserOK()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = lh.Persister

	if err := librarian.RemoveUser(params.User); err != nil {
		log.Println(err)
		return operations.NewDeleteLibrarianUsernameUserUserOK()
	}

	return operations.NewDeleteLibrarianUsernameUserUserOK()
}

// HandleGetUser is the handler for get request for the endpoint /librarian/{username}/user/{user}
func (lh *LibrarianHandler) HandleGetUser(params operations.GetLibrarianUsernameUserUserParams) middleware.Responder {

	// get the librarian
	filter := bson.D{{"username", params.Username}}

	var librarian library.Librarian
	if err := lh.Persister.Retrieve("librarians", filter, &librarian); err != nil {
		log.Println(err)
		return operations.NewGetLibrarianUsernameUserUserOK()
	}
	
	// ensure the librarian has a valid persister
	librarian.Persister = lh.Persister

	user, err := librarian.User(params.User)
	if err != nil {
		log.Println(err)
		return operations.NewGetLibrarianUsernameUserUserOK()
	}

	responseUser := &models.User{
		Name: user.Name,
		Username: user.Username,
	}

	return operations.NewGetLibrarianUsernameUserUserOK().WithPayload(responseUser)
}