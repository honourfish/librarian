package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockPersister is a mock persister for the Persister interface.
type MockPersister struct {
	mock.Mock
}

func (mp *MockPersister) Create(collection string, object interface{}) error {
	args := mp.Called(collection, object)
	return args.Error(0)
}

func (mp *MockPersister) Update(collection string, filter interface{}, object interface{}) error {
	args := mp.Called(collection, filter, object)
	return args.Error(0)
}

func (mp *MockPersister) Retrieve(collection string, filter interface{}, result interface{}) error {
	args := mp.Called(collection, filter, result)
	return args.Error(0)
}

func (mp *MockPersister) Delete(collection string, filter interface{}) error {
	args := mp.Called(collection, filter)
	return args.Error(0)
}