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
)

// NewPutLibrarianUsernameUserUserCheckoutParams creates a new PutLibrarianUsernameUserUserCheckoutParams object
//
// There are no default values defined in the spec.
func NewPutLibrarianUsernameUserUserCheckoutParams() PutLibrarianUsernameUserUserCheckoutParams {

	return PutLibrarianUsernameUserUserCheckoutParams{}
}

// PutLibrarianUsernameUserUserCheckoutParams contains all the bound params for the put librarian username user user checkout operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutLibrarianUsernameUserUserCheckout
type PutLibrarianUsernameUserUserCheckoutParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Book PutLibrarianUsernameUserUserCheckoutBody
	/*
	  Required: true
	  In: path
	*/
	User string
	/*
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutLibrarianUsernameUserUserCheckoutParams() beforehand.
func (o *PutLibrarianUsernameUserUserCheckoutParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PutLibrarianUsernameUserUserCheckoutBody
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
				o.Book = body
			}
		}
	}

	rUser, rhkUser, _ := route.Params.GetOK("user")
	if err := o.bindUser(rUser, rhkUser, route.Formats); err != nil {
		res = append(res, err)
	}

	rUsername, rhkUsername, _ := route.Params.GetOK("username")
	if err := o.bindUsername(rUsername, rhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUser binds and validates parameter User from path.
func (o *PutLibrarianUsernameUserUserCheckoutParams) bindUser(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.User = raw

	return nil
}

// bindUsername binds and validates parameter Username from path.
func (o *PutLibrarianUsernameUserUserCheckoutParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	return nil
}
