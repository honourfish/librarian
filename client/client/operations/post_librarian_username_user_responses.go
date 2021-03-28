// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostLibrarianUsernameUserReader is a Reader for the PostLibrarianUsernameUser structure.
type PostLibrarianUsernameUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLibrarianUsernameUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLibrarianUsernameUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPostLibrarianUsernameUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLibrarianUsernameUserCreated creates a PostLibrarianUsernameUserCreated with default headers values
func NewPostLibrarianUsernameUserCreated() *PostLibrarianUsernameUserCreated {
	return &PostLibrarianUsernameUserCreated{}
}

/* PostLibrarianUsernameUserCreated describes a response with status code 201, with default header values.

Created
*/
type PostLibrarianUsernameUserCreated struct {
}

func (o *PostLibrarianUsernameUserCreated) Error() string {
	return fmt.Sprintf("[POST /librarian/{username}/user][%d] postLibrarianUsernameUserCreated ", 201)
}

func (o *PostLibrarianUsernameUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLibrarianUsernameUserNotFound creates a PostLibrarianUsernameUserNotFound with default headers values
func NewPostLibrarianUsernameUserNotFound() *PostLibrarianUsernameUserNotFound {
	return &PostLibrarianUsernameUserNotFound{}
}

/* PostLibrarianUsernameUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostLibrarianUsernameUserNotFound struct {
}

func (o *PostLibrarianUsernameUserNotFound) Error() string {
	return fmt.Sprintf("[POST /librarian/{username}/user][%d] postLibrarianUsernameUserNotFound ", 404)
}

func (o *PostLibrarianUsernameUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
