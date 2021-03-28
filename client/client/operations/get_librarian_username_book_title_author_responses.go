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

// GetLibrarianUsernameBookTitleAuthorReader is a Reader for the GetLibrarianUsernameBookTitleAuthor structure.
type GetLibrarianUsernameBookTitleAuthorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLibrarianUsernameBookTitleAuthorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLibrarianUsernameBookTitleAuthorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLibrarianUsernameBookTitleAuthorOK creates a GetLibrarianUsernameBookTitleAuthorOK with default headers values
func NewGetLibrarianUsernameBookTitleAuthorOK() *GetLibrarianUsernameBookTitleAuthorOK {
	return &GetLibrarianUsernameBookTitleAuthorOK{}
}

/* GetLibrarianUsernameBookTitleAuthorOK describes a response with status code 200, with default header values.

Retrieved
*/
type GetLibrarianUsernameBookTitleAuthorOK struct {
	Payload *models.BookStock
}

func (o *GetLibrarianUsernameBookTitleAuthorOK) Error() string {
	return fmt.Sprintf("[GET /librarian/{username}/book/{title}/{author}][%d] getLibrarianUsernameBookTitleAuthorOK  %+v", 200, o.Payload)
}
func (o *GetLibrarianUsernameBookTitleAuthorOK) GetPayload() *models.BookStock {
	return o.Payload
}

func (o *GetLibrarianUsernameBookTitleAuthorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BookStock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}