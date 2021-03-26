package library

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoDB struct {
	// Timeout is the time (in seconds) to wait for the connection to the mongo server.
	Timeout int

	// URI is the mongodb uri
	URI string

	// Database is the database to add persistant data to.
	Database string
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

// Create adds an object to a collection
func (mdb *MongoDB) Create(collection string, object interface{}) error {
	// make a connection to the database
	client, ctx, cancel, err := mdb.getContext()
	if err != nil {
		return err
	}
	defer cancel()
	defer client.Disconnect(ctx)

	database := client.Database(mdb.Database)
	dbCollection := database.Collection(collection)

	_, err = dbCollection.InsertOne(ctx, object)
	return err
}

// Retrieve gets a single object from a collection, given a specific search filter.
// The given filter can be a bson.D object.
// Returns a *mongo.SingleResult, which can use `Decode(interface{})` to decode the result into a bson
//   compatible object.
func (mdb *MongoDB) Retrieve(collection string, filter interface{}) (interface{}, error) {
	// make a connection to the database
	client, ctx, cancel, err := mdb.getContext()
	if err != nil {
		return nil, err
	}
	defer cancel()
	defer client.Disconnect(ctx)

	database := client.Database(mdb.Database)
	dbCollection := database.Collection(collection)

	var object *mongo.SingleResult
	object = dbCollection.FindOne(ctx, filter)
	return object, err
}

// Update gets a single object from a collection, based on a given filter, then replaces it
//   with a given object.
// The given filter can be a bson.D object.
func (mdb *MongoDB) Update(collection string, filter interface{}, object interface{}) error {
	// make a connection to the database
	client, ctx, cancel, err := mdb.getContext()
	if err != nil {
		return err
	}
	defer cancel()
	defer client.Disconnect(ctx)

	database := client.Database(mdb.Database)
	dbCollection := database.Collection(collection)

	_, err = dbCollection.ReplaceOne(ctx, filter, object)
	return err
}

// Delete removes a single object from a collection, based on a given filter.
// The given filter can be a bson.D object.
func (mdb *MongoDB) Delete(collection string, filter interface{}) error {
	// make a connection to the database
	client, ctx, cancel, err := mdb.getContext()
	if err != nil {
		return err
	}
	defer cancel()
	defer client.Disconnect(ctx)

	database := client.Database(mdb.Database)
	dbCollection := database.Collection(collection)

	_, err = dbCollection.DeleteOne(ctx, filter)
	return err
}