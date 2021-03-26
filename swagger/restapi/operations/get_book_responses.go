// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"librarian/swagger/models"
)

// GetBookOKCode is the HTTP code returned for type GetBookOK
const GetBookOKCode int = 200

/*GetBookOK Retrieved

swagger:response getBookOK
*/
type GetBookOK struct {

	/*
	  In: Body
	*/
	Payload *models.Book `json:"body,omitempty"`
}

// NewGetBookOK creates GetBookOK with default headers values
func NewGetBookOK() *GetBookOK {

	return &GetBookOK{}
}

// WithPayload adds the payload to the get book o k response
func (o *GetBookOK) WithPayload(payload *models.Book) *GetBookOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get book o k response
func (o *GetBookOK) SetPayload(payload *models.Book) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBookOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
