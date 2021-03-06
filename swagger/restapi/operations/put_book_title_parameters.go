// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"librarian/swagger/models"
)

// NewPutBookTitleParams creates a new PutBookTitleParams object
//
// There are no default values defined in the spec.
func NewPutBookTitleParams() PutBookTitleParams {

	return PutBookTitleParams{}
}

// PutBookTitleParams contains all the bound params for the put book title operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutBookTitle
type PutBookTitleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Book *models.Book
	/*
	  Required: true
	  In: path
	*/
	Title string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutBookTitleParams() beforehand.
func (o *PutBookTitleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Book
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("book", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Book = &body
			}
		}
	}

	rTitle, rhkTitle, _ := route.Params.GetOK("title")
	if err := o.bindTitle(rTitle, rhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTitle binds and validates parameter Title from path.
func (o *PutBookTitleParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Title = raw

	return nil
}
