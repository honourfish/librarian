// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetBookParams creates a new GetBookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBookParams() *GetBookParams {
	return &GetBookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBookParamsWithTimeout creates a new GetBookParams object
// with the ability to set a timeout on a request.
func NewGetBookParamsWithTimeout(timeout time.Duration) *GetBookParams {
	return &GetBookParams{
		timeout: timeout,
	}
}

// NewGetBookParamsWithContext creates a new GetBookParams object
// with the ability to set a context for a request.
func NewGetBookParamsWithContext(ctx context.Context) *GetBookParams {
	return &GetBookParams{
		Context: ctx,
	}
}

// NewGetBookParamsWithHTTPClient creates a new GetBookParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBookParamsWithHTTPClient(client *http.Client) *GetBookParams {
	return &GetBookParams{
		HTTPClient: client,
	}
}

/* GetBookParams contains all the parameters to send to the API endpoint
   for the get book operation.

   Typically these are written to a http.Request.
*/
type GetBookParams struct {

	// Title.
	Title string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get book params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBookParams) WithDefaults() *GetBookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get book params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get book params
func (o *GetBookParams) WithTimeout(timeout time.Duration) *GetBookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get book params
func (o *GetBookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get book params
func (o *GetBookParams) WithContext(ctx context.Context) *GetBookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get book params
func (o *GetBookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get book params
func (o *GetBookParams) WithHTTPClient(client *http.Client) *GetBookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get book params
func (o *GetBookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTitle adds the title to the get book params
func (o *GetBookParams) WithTitle(title string) *GetBookParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the get book params
func (o *GetBookParams) SetTitle(title string) {
	o.Title = title
}

// WriteToRequest writes these params to a swagger request
func (o *GetBookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param title
	qrTitle := o.Title
	qTitle := qrTitle
	if qTitle != "" {

		if err := r.SetQueryParam("title", qTitle); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
