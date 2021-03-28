package handlers

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/go-openapi/runtime/middleware"

	"librarian/swagger/restapi/operations"
	"librarian/swagger/models"

	"librarian/library/data"
	"librarian/library/interfaces"
)

type BookHandler struct {
	Persister interfaces.Persister
}


// handlePostBook is the handler for post request for the endpoint /book
func (bh *BookHandler) HandlePostBook(params operations.PostBookParams) middleware.Responder {
	book := &data.Book{
		Title: params.Book.Title,
		Author: params.Book.Author,
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

	var book data.Book
	if err := bh.Persister.Retrieve("books", filter, &book); err != nil {
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

// handleGetBookTitle is the handler for get request for the endpoint /book/{title}
func (bh *BookHandler) HandleGetBookTitle(params operations.GetBookTitleParams) middleware.Responder {

	filter := bson.D{{"title", params.Title}}

	log.Println(params.Title)

	var book data.Book
	if err := bh.Persister.Retrieve("books", filter, &book); err != nil {
		log.Println(err)
		return operations.NewGetBookTitleOK()
	}

	responseBook := &models.Book{
		Title: book.Title,
		Author: book.Author,
	}

	log.Println(responseBook.Title)
	return operations.NewGetBookTitleOK().WithPayload(responseBook)
}

// handlePutBookTitle is the handler for put request for the endpoint /book/{title}
func (bh *BookHandler) HandlePutBookTitle(params operations.PutBookTitleParams) middleware.Responder {

	filter := bson.D{{"title", params.Title}}

	book := &data.Book{
		Title: params.Book.Title,
		Author: params.Book.Author,
		PublishDate: time.Now(),
	}

	log.Println(params.Title)

	if err := bh.Persister.Update("books", filter, book); err != nil {
		log.Println(err)
		return operations.NewPutBookTitleOK()
	}

	return operations.NewPutBookTitleOK()
}

// handleDeleteBookTitle is the handler for delete request for the endpoint /book/{title}
func (bh *BookHandler) HandleDeleteBookTitle(params operations.DeleteBookTitleParams) middleware.Responder {

	filter := bson.D{{"title", params.Title}}

	if err := bh.Persister.Delete("books", filter); err != nil {
		log.Println(err)
		return operations.NewDeleteBookTitleOK()
	}

	return operations.NewDeleteBookTitleOK()
}