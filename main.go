package main

import (
	//"fmt"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/readpref"

	"librarian/librarian"
	"librarian/librarian/interfaces"
)


type FooBar struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Foo string `bson:"foo,omitempty"`
}

func main() {

	mongoDb := &librarian.MongoDB{
		Timeout: 10,
		URI: "mongodb://root:example@localhost:27017",	
	}

	var persist persister.Persister = mongoDb
	var bookPersister persister.BookPersister = mongoDb

	err := persist.InitPersister()
	if err != nil {
		panic(err)
	}
	defer persist.ClosePersister()

	book := librarian.Book{
		Title: "Moby Prejedice",
		Author: "bebop",
		PublishDate: time.Now(),
	}

	err = bookPersister.AddBook(book)

	if err != nil {
		panic(err)
	}

}