package handlers

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-openapi/runtime/middleware"

	"librarian/swagger/restapi/operations"
	"librarian/swagger/models"

	"librarian/library"
	"librarian/library/interfaces"
)

type BookHandler struct {
	Persister persistance.Persister
}


// handlePostBook is the handler for post request for the endpoint /book
func (bh *BookHandler) HandlePostBook(params operations.PostBookParams) middleware.Responder {
	book := &library.Book{
		Title: params.Book.Title,
		Author: params.Book.Author,
		PublishDate: time.Now(),
	}

	if err := bh.Persister.Create("books", book); err != nil {
		log.Println(err)
	}

	log.Println(params.Book.Title)

	return operations.NewPostBookCreated()
}

// handleGetBook is the handler for get request for the endpoint /book
func (bh *BookHandler) HandleGetBook(params operations.GetBookParams) middleware.Responder {

	filter := bson.D{{"title", params.Title}}

	log.Println(params.Title)

	result, err := bh.Persister.Retrieve("books", filter)
	if err != nil {
		log.Println(err)
		return operations.NewGetBookOK()
	}

	var book library.Book
	if err = result.(*mongo.SingleResult).Decode(&book); err != nil {
		log.Println(err)
		return operations.NewGetBookOK()
	}

	responseBook := &models.Book{
		Title: book.Title,
		Author: book.Author,
	}

	log.Println(responseBook.Title)
	return operations.NewGetBookOK().WithPayload(responseBook)
}