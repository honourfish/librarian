// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetBookParams creates a new GetBookParams object
//
// There are no default values defined in the spec.
func NewGetBookParams() GetBookParams {

	return GetBookParams{}
}

// GetBookParams contains all the bound params for the get book operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetBook
type GetBookParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	Title string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBookParams() beforehand.
func (o *GetBookParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qTitle, qhkTitle, _ := qs.GetOK("title")
	if err := o.bindTitle(qTitle, qhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTitle binds and validates parameter Title from query.
func (o *GetBookParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("title", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("title", "query", raw); err != nil {
		return err
	}
	o.Title = raw

	return nil
}
