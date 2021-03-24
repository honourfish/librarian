package librarian

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

	// client is the mongo db client
	client *mongo.Client

	// ctx is the mongodb connection context
	ctx context.Context

	// cancel is the cancel function to call when we cancel the context
	cancel context.CancelFunc
}

// InitPersister creates a new mongo db client and connection.
func (mdb *MongoDB) InitPersister() (err error) {

	mdb.client, err = mongo.NewClient(options.Client().ApplyURI(mdb.URI))
	if err != nil {
		return
	}

	mdb.ctx, mdb.cancel = context.WithTimeout(context.Background(), time.Duration(mdb.Timeout)*time.Second)

	err = mdb.client.Connect(mdb.ctx)
	if err != nil {
		return
	}

	return
}

// ClosePersister closes the mongodb
func (mdb *MongoDB) ClosePersister() (err error) {
	mdb.client.Disconnect(mdb.ctx)
	mdb.cancel()
	return
}

// AddBook adds a book
func (mdb *MongoDB) AddBook(book Book) (err error) {
	database := mdb.client.Database("library")
	booksCollection := database.Collection("books")

	_, err = booksCollection.InsertOne(mdb.ctx, book)
	return
}

func (mdb *MongoDB) GetBook(id primitive.ObjectID) (book Book, err error) {
	database := mdb.client.Database("library")
	booksCollection := database.Collection("books")

	filter := bson.D{{"ID", id}}

	err = booksCollection.FindOne(mdb.ctx, filter).Decode(&book)
	return
}