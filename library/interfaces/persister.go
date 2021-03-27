package interfaces

type Persister interface {
	// Create persists a given object within a given collection.
	Create(collection string, object interface{}) error

	// Retrieve gets an object from a given collection, given a filter.
	Retrieve(collection string, filter interface{}, result interface{}) error

	// Update updates an object from a given collection, given a filter.
	//   and the updated object.
	Update(collection string, filter interface{}, object interface{}) error

	// Delete deletes an object from a given collection, given a filter.
	Delete(collection string, filter interface{}) error
}