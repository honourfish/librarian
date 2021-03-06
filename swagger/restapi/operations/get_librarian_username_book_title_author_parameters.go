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

// NewGetLibrarianUsernameBookTitleAuthorParams creates a new GetLibrarianUsernameBookTitleAuthorParams object
//
// There are no default values defined in the spec.
func NewGetLibrarianUsernameBookTitleAuthorParams() GetLibrarianUsernameBookTitleAuthorParams {

	return GetLibrarianUsernameBookTitleAuthorParams{}
}

// GetLibrarianUsernameBookTitleAuthorParams contains all the bound params for the get librarian username book title author operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetLibrarianUsernameBookTitleAuthor
type GetLibrarianUsernameBookTitleAuthorParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Author string
	/*
	  Required: true
	  In: path
	*/
	Title string
	/*
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLibrarianUsernameBookTitleAuthorParams() beforehand.
func (o *GetLibrarianUsernameBookTitleAuthorParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAuthor, rhkAuthor, _ := route.Params.GetOK("author")
	if err := o.bindAuthor(rAuthor, rhkAuthor, route.Formats); err != nil {
		res = append(res, err)
	}

	rTitle, rhkTitle, _ := route.Params.GetOK("title")
	if err := o.bindTitle(rTitle, rhkTitle, route.Formats); err != nil {
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

// bindAuthor binds and validates parameter Author from path.
func (o *GetLibrarianUsernameBookTitleAuthorParams) bindAuthor(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Author = raw

	return nil
}

// bindTitle binds and validates parameter Title from path.
func (o *GetLibrarianUsernameBookTitleAuthorParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Title = raw

	return nil
}

// bindUsername binds and validates parameter Username from path.
func (o *GetLibrarianUsernameBookTitleAuthorParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	return nil
}
