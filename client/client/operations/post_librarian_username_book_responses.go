// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostLibrarianUsernameBookReader is a Reader for the PostLibrarianUsernameBook structure.
type PostLibrarianUsernameBookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLibrarianUsernameBookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLibrarianUsernameBookCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLibrarianUsernameBookCreated creates a PostLibrarianUsernameBookCreated with default headers values
func NewPostLibrarianUsernameBookCreated() *PostLibrarianUsernameBookCreated {
	return &PostLibrarianUsernameBookCreated{}
}

/* PostLibrarianUsernameBookCreated describes a response with status code 201, with default header values.

Created
*/
type PostLibrarianUsernameBookCreated struct {
}

func (o *PostLibrarianUsernameBookCreated) Error() string {
	return fmt.Sprintf("[POST /librarian/{username}/book][%d] postLibrarianUsernameBookCreated ", 201)
}

func (o *PostLibrarianUsernameBookCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
