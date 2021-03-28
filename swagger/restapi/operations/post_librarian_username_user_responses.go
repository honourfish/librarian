// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostLibrarianUsernameUserCreatedCode is the HTTP code returned for type PostLibrarianUsernameUserCreated
const PostLibrarianUsernameUserCreatedCode int = 201

/*PostLibrarianUsernameUserCreated Created

swagger:response postLibrarianUsernameUserCreated
*/
type PostLibrarianUsernameUserCreated struct {
}

// NewPostLibrarianUsernameUserCreated creates PostLibrarianUsernameUserCreated with default headers values
func NewPostLibrarianUsernameUserCreated() *PostLibrarianUsernameUserCreated {

	return &PostLibrarianUsernameUserCreated{}
}

// WriteResponse to the client
func (o *PostLibrarianUsernameUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}
