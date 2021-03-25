package main

import (
	//"fmt"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/readpref"


	// swagger stuff
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	//"github.com/go-openapi/swag"

	//"librarian/swagger/models"
	"librarian/swagger/restapi"
	"librarian/swagger/restapi/operations"

	"librarian/library"
	"librarian/handlers"
)


type FooBar struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Foo string `bson:"foo,omitempty"`
}

// handlePostBook is the handler for post request for the endpoint /book
func handlePostBook(params operations.PostBookParams) middleware.Responder {
	log.Println(params.Book.Title)
	return operations.NewPostBookCreated()
}

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