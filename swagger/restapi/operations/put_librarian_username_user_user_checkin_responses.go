// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutLibrarianUsernameUserUserCheckinOKCode is the HTTP code returned for type PutLibrarianUsernameUserUserCheckinOK
const PutLibrarianUsernameUserUserCheckinOKCode int = 200

/*PutLibrarianUsernameUserUserCheckinOK Updated

swagger:response putLibrarianUsernameUserUserCheckinOK
*/
type PutLibrarianUsernameUserUserCheckinOK struct {
}

// NewPutLibrarianUsernameUserUserCheckinOK creates PutLibrarianUsernameUserUserCheckinOK with default headers values
func NewPutLibrarianUsernameUserUserCheckinOK() *PutLibrarianUsernameUserUserCheckinOK {

	return &PutLibrarianUsernameUserUserCheckinOK{}
}

// WriteResponse to the client
func (o *PutLibrarianUsernameUserUserCheckinOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutLibrarianUsernameUserUserCheckinNotFoundCode is the HTTP code returned for type PutLibrarianUsernameUserUserCheckinNotFound
const PutLibrarianUsernameUserUserCheckinNotFoundCode int = 404

/*PutLibrarianUsernameUserUserCheckinNotFound Not Found

swagger:response putLibrarianUsernameUserUserCheckinNotFound
*/
type PutLibrarianUsernameUserUserCheckinNotFound struct {
}

// NewPutLibrarianUsernameUserUserCheckinNotFound creates PutLibrarianUsernameUserUserCheckinNotFound with default headers values
func NewPutLibrarianUsernameUserUserCheckinNotFound() *PutLibrarianUsernameUserUserCheckinNotFound {

	return &PutLibrarianUsernameUserUserCheckinNotFound{}
}

// WriteResponse to the client
func (o *PutLibrarianUsernameUserUserCheckinNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
