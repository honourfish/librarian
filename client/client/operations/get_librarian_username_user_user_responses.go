// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"librarian/client/models"
)

// GetLibrarianUsernameUserUserReader is a Reader for the GetLibrarianUsernameUserUser structure.
type GetLibrarianUsernameUserUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLibrarianUsernameUserUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLibrarianUsernameUserUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLibrarianUsernameUserUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLibrarianUsernameUserUserOK creates a GetLibrarianUsernameUserUserOK with default headers values
func NewGetLibrarianUsernameUserUserOK() *GetLibrarianUsernameUserUserOK {
	return &GetLibrarianUsernameUserUserOK{}
}

/* GetLibrarianUsernameUserUserOK describes a response with status code 200, with default header values.

Retrieved
*/
type GetLibrarianUsernameUserUserOK struct {
	Payload *models.User
}

func (o *GetLibrarianUsernameUserUserOK) Error() string {
	return fmt.Sprintf("[GET /librarian/{username}/user/{user}][%d] getLibrarianUsernameUserUserOK  %+v", 200, o.Payload)
}
func (o *GetLibrarianUsernameUserUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetLibrarianUsernameUserUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibrarianUsernameUserUserNotFound creates a GetLibrarianUsernameUserUserNotFound with default headers values
func NewGetLibrarianUsernameUserUserNotFound() *GetLibrarianUsernameUserUserNotFound {
	return &GetLibrarianUsernameUserUserNotFound{}
}

/* GetLibrarianUsernameUserUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetLibrarianUsernameUserUserNotFound struct {
}

func (o *GetLibrarianUsernameUserUserNotFound) Error() string {
	return fmt.Sprintf("[GET /librarian/{username}/user/{user}][%d] getLibrarianUsernameUserUserNotFound ", 404)
}

func (o *GetLibrarianUsernameUserUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
