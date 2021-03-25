package handlers

import (
	"log"
	"time"
	"github.com/go-openapi/runtime/middleware"

	"librarian/swagger/restapi/operations"

	"librarian/library"
	"librarian/library/interfaces"
)

type BookHandler struct {
	BookPersister persister.BookPersister
}


// handlePostBook is the handler for post request for the endpoint /book
func (bh *BookHandler) HandlePostBook(params operations.PostBookParams) middleware.Responder {
	book := library.Book{
		Title: params.Book.Title,
		Author: params.Book.Author,
		PublishDate: time.Now(),
	}

	if err := bh.BookPersister.AddBook(book); err != nil {
		log.Println(err)
	}

	log.Println(params.Book.Title)

	return operations.NewPostBookCreated()
}