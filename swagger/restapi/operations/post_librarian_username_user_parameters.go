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

// NewPostLibrarianUsernameUserParams creates a new PostLibrarianUsernameUserParams object
//
// There are no default values defined in the spec.
func NewPostLibrarianUsernameUserParams() PostLibrarianUsernameUserParams {

	return PostLibrarianUsernameUserParams{}
}

// PostLibrarianUsernameUserParams contains all the bound params for the post librarian username user operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostLibrarianUsernameUser
type PostLibrarianUsernameUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	User *models.User
	/*
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostLibrarianUsernameUserParams() beforehand.
func (o *PostLibrarianUsernameUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.User
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("user", "body", "", err))
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
				o.User = &body
			}
		}
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

// bindUsername binds and validates parameter Username from path.
func (o *PostLibrarianUsernameUserParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	return nil
}
