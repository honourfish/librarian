// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutLibrarianUsernameUserUserCheckoutOKCode is the HTTP code returned for type PutLibrarianUsernameUserUserCheckoutOK
const PutLibrarianUsernameUserUserCheckoutOKCode int = 200

/*PutLibrarianUsernameUserUserCheckoutOK Updated

swagger:response putLibrarianUsernameUserUserCheckoutOK
*/
type PutLibrarianUsernameUserUserCheckoutOK struct {
}

// NewPutLibrarianUsernameUserUserCheckoutOK creates PutLibrarianUsernameUserUserCheckoutOK with default headers values
func NewPutLibrarianUsernameUserUserCheckoutOK() *PutLibrarianUsernameUserUserCheckoutOK {

	return &PutLibrarianUsernameUserUserCheckoutOK{}
}

// WriteResponse to the client
func (o *PutLibrarianUsernameUserUserCheckoutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutLibrarianUsernameUserUserCheckoutNotFoundCode is the HTTP code returned for type PutLibrarianUsernameUserUserCheckoutNotFound
const PutLibrarianUsernameUserUserCheckoutNotFoundCode int = 404

/*PutLibrarianUsernameUserUserCheckoutNotFound Not Found

swagger:response putLibrarianUsernameUserUserCheckoutNotFound
*/
type PutLibrarianUsernameUserUserCheckoutNotFound struct {
}

// NewPutLibrarianUsernameUserUserCheckoutNotFound creates PutLibrarianUsernameUserUserCheckoutNotFound with default headers values
func NewPutLibrarianUsernameUserUserCheckoutNotFound() *PutLibrarianUsernameUserUserCheckoutNotFound {

	return &PutLibrarianUsernameUserUserCheckoutNotFound{}
}

// WriteResponse to the client
func (o *PutLibrarianUsernameUserUserCheckoutNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}