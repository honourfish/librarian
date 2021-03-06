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

// NewDeleteBookTitleParams creates a new DeleteBookTitleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBookTitleParams() *DeleteBookTitleParams {
	return &DeleteBookTitleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBookTitleParamsWithTimeout creates a new DeleteBookTitleParams object
// with the ability to set a timeout on a request.
func NewDeleteBookTitleParamsWithTimeout(timeout time.Duration) *DeleteBookTitleParams {
	return &DeleteBookTitleParams{
		timeout: timeout,
	}
}

// NewDeleteBookTitleParamsWithContext creates a new DeleteBookTitleParams object
// with the ability to set a context for a request.
func NewDeleteBookTitleParamsWithContext(ctx context.Context) *DeleteBookTitleParams {
	return &DeleteBookTitleParams{
		Context: ctx,
	}
}

// NewDeleteBookTitleParamsWithHTTPClient creates a new DeleteBookTitleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBookTitleParamsWithHTTPClient(client *http.Client) *DeleteBookTitleParams {
	return &DeleteBookTitleParams{
		HTTPClient: client,
	}
}

/* DeleteBookTitleParams contains all the parameters to send to the API endpoint
   for the delete book title operation.

   Typically these are written to a http.Request.
*/
type DeleteBookTitleParams struct {

	// Title.
	Title string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete book title params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBookTitleParams) WithDefaults() *DeleteBookTitleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete book title params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBookTitleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete book title params
func (o *DeleteBookTitleParams) WithTimeout(timeout time.Duration) *DeleteBookTitleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete book title params
func (o *DeleteBookTitleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete book title params
func (o *DeleteBookTitleParams) WithContext(ctx context.Context) *DeleteBookTitleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete book title params
func (o *DeleteBookTitleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete book title params
func (o *DeleteBookTitleParams) WithHTTPClient(client *http.Client) *DeleteBookTitleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete book title params
func (o *DeleteBookTitleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTitle adds the title to the delete book title params
func (o *DeleteBookTitleParams) WithTitle(title string) *DeleteBookTitleParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the delete book title params
func (o *DeleteBookTitleParams) SetTitle(title string) {
	o.Title = title
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBookTitleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param title
	if err := r.SetPathParam("title", o.Title); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
