// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostLibrarianUsernameBookCreatedCode is the HTTP code returned for type PostLibrarianUsernameBookCreated
const PostLibrarianUsernameBookCreatedCode int = 201

/*PostLibrarianUsernameBookCreated Created

swagger:response postLibrarianUsernameBookCreated
*/
type PostLibrarianUsernameBookCreated struct {
}

// NewPostLibrarianUsernameBookCreated creates PostLibrarianUsernameBookCreated with default headers values
func NewPostLibrarianUsernameBookCreated() *PostLibrarianUsernameBookCreated {

	return &PostLibrarianUsernameBookCreated{}
}

// WriteResponse to the client
func (o *PostLibrarianUsernameBookCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}