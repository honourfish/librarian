// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"librarian/swagger/models"
)

// GetLibrarianUsernameBookTitleAuthorOKCode is the HTTP code returned for type GetLibrarianUsernameBookTitleAuthorOK
const GetLibrarianUsernameBookTitleAuthorOKCode int = 200

/*GetLibrarianUsernameBookTitleAuthorOK Retrieved

swagger:response getLibrarianUsernameBookTitleAuthorOK
*/
type GetLibrarianUsernameBookTitleAuthorOK struct {

	/*
	  In: Body
	*/
	Payload *models.BookStock `json:"body,omitempty"`
}

// NewGetLibrarianUsernameBookTitleAuthorOK creates GetLibrarianUsernameBookTitleAuthorOK with default headers values
func NewGetLibrarianUsernameBookTitleAuthorOK() *GetLibrarianUsernameBookTitleAuthorOK {

	return &GetLibrarianUsernameBookTitleAuthorOK{}
}

// WithPayload adds the payload to the get librarian username book title author o k response
func (o *GetLibrarianUsernameBookTitleAuthorOK) WithPayload(payload *models.BookStock) *GetLibrarianUsernameBookTitleAuthorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get librarian username book title author o k response
func (o *GetLibrarianUsernameBookTitleAuthorOK) SetPayload(payload *models.BookStock) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLibrarianUsernameBookTitleAuthorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}