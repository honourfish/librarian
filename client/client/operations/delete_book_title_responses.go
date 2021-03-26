// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteBookTitleReader is a Reader for the DeleteBookTitle structure.
type DeleteBookTitleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBookTitleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBookTitleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteBookTitleOK creates a DeleteBookTitleOK with default headers values
func NewDeleteBookTitleOK() *DeleteBookTitleOK {
	return &DeleteBookTitleOK{}
}

/* DeleteBookTitleOK describes a response with status code 200, with default header values.

Deleted
*/
type DeleteBookTitleOK struct {
}

func (o *DeleteBookTitleOK) Error() string {
	return fmt.Sprintf("[DELETE /book/{title}][%d] deleteBookTitleOK ", 200)
}

func (o *DeleteBookTitleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
