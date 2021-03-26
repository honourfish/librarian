// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostBookReader is a Reader for the PostBook structure.
type PostBookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostBookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostBookCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostBookCreated creates a PostBookCreated with default headers values
func NewPostBookCreated() *PostBookCreated {
	return &PostBookCreated{}
}

/* PostBookCreated describes a response with status code 201, with default header values.

Created
*/
type PostBookCreated struct {
}

func (o *PostBookCreated) Error() string {
	return fmt.Sprintf("[POST /book][%d] postBookCreated ", 201)
}

func (o *PostBookCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
