// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteLibrarianUsernameUserUserUsernameParams creates a new DeleteLibrarianUsernameUserUserUsernameParams object
//
// There are no default values defined in the spec.
func NewDeleteLibrarianUsernameUserUserUsernameParams() DeleteLibrarianUsernameUserUserUsernameParams {

	return DeleteLibrarianUsernameUserUserUsernameParams{}
}

// DeleteLibrarianUsernameUserUserUsernameParams contains all the bound params for the delete librarian username user user username operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteLibrarianUsernameUserUserUsername
type DeleteLibrarianUsernameUserUserUsernameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	UserUsername string
	/*
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteLibrarianUsernameUserUserUsernameParams() beforehand.
func (o *DeleteLibrarianUsernameUserUserUsernameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUserUsername, rhkUserUsername, _ := route.Params.GetOK("user_username")
	if err := o.bindUserUsername(rUserUsername, rhkUserUsername, route.Formats); err != nil {
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

// bindUserUsername binds and validates parameter UserUsername from path.
func (o *DeleteLibrarianUsernameUserUserUsernameParams) bindUserUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UserUsername = raw

	return nil
}

// bindUsername binds and validates parameter Username from path.
func (o *DeleteLibrarianUsernameUserUserUsernameParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	return nil
}