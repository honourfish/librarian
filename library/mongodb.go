package library

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoDB struct {
	// Timeout is the time (in seconds) to wait for the connection to the mongo server.
	Timeout int

	// URI is the mongodb uri
	URI string
}

// getContext gets a client and context.
func (mdb *MongoDB) getContext() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc, err error) {
	client, err = mongo.NewClient(options.Client().ApplyURI(mdb.URI))
	if err != nil {
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(mdb.Timeout)*time.Second)

	err = client.Connect(ctx)
	return
}

// AddBook adds a book
func (mdb *MongoDB) AddBook(book Book) error {
	// make a connection to the database
	client, ctx, cancel, err := mdb.getContext()
	defer cancel()
	defer client.Disconnect(ctx)

	database := client.Database("library")
	booksCollection := database.Collection("books")

	_, err = booksCollection.InsertOne(ctx, book)
	return err
}

func (mdb *MongoDB) GetBook(id primitive.ObjectID) (Book, error) {
	// make a connection to the database
	client, ctx, cancel, err := mdb.getContext()
	defer cancel()
	defer client.Disconnect(ctx)

	database := client.Database("library")
	booksCollection := database.Collection("books")

	filter := bson.D{{"ID", id}}

	var book Book
	err = booksCollection.FindOne(ctx, filter).Decode(&book)
	return book, err
}