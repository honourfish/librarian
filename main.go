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

	server.Port = 8082

	bookHandler := handlers.BookHandler{BookPersister: mongoDb}
	api.PostBookHandler = operations.PostBookHandlerFunc(bookHandler.HandlePostBook)

	if err = server.Serve(); err != nil {
		log.Fatalln(err)
	}
}