package main

import (
	//"fmt"
	"log"


	// swagger stuff
	"github.com/go-openapi/loads"
	
	"librarian/swagger/restapi"
	"librarian/swagger/restapi/operations"

	"librarian/library"
	"librarian/handlers"
)

func main() {

	mongoDb := &library.MongoDB{
		Timeout: 3,
		URI: "mongodb://root:example@localhost:27017",	
		Database: "library",
	}

	// do the server stuff
	var swaggerSpec *loads.Document
	var err error
	if swaggerSpec, err = loads.Analyzed(restapi.SwaggerJSON, ""); err != nil {
		log.Fatalln(err)
	}

	// create a new API which can be used to create a server
	api := operations.NewLibraryAPI(swaggerSpec)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080

	bookHandler := handlers.BookHandler{Persister: mongoDb}

	// /book handlers
	api.PostBookHandler = operations.PostBookHandlerFunc(bookHandler.HandlePostBook)
	api.GetBookHandler = operations.GetBookHandlerFunc(bookHandler.HandleGetBook)

	// /book/{title} handlers
	api.GetBookTitleHandler = operations.GetBookTitleHandlerFunc(bookHandler.HandleGetBookTitle)
	api.PutBookTitleHandler = operations.PutBookTitleHandlerFunc(bookHandler.HandlePutBookTitle)
	api.DeleteBookTitleHandler = operations.DeleteBookTitleHandlerFunc(bookHandler.HandleDeleteBookTitle)

	if err = server.Serve(); err != nil {
		log.Fatalln(err)
	}
}