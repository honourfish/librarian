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

	"librarian/client/models"
)

// NewPostLibrarianUsernameBookParams creates a new PostLibrarianUsernameBookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLibrarianUsernameBookParams() *PostLibrarianUsernameBookParams {
	return &PostLibrarianUsernameBookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLibrarianUsernameBookParamsWithTimeout creates a new PostLibrarianUsernameBookParams object
// with the ability to set a timeout on a request.
func NewPostLibrarianUsernameBookParamsWithTimeout(timeout time.Duration) *PostLibrarianUsernameBookParams {
	return &PostLibrarianUsernameBookParams{
		timeout: timeout,
	}
}

// NewPostLibrarianUsernameBookParamsWithContext creates a new PostLibrarianUsernameBookParams object
// with the ability to set a context for a request.
func NewPostLibrarianUsernameBookParamsWithContext(ctx context.Context) *PostLibrarianUsernameBookParams {
	return &PostLibrarianUsernameBookParams{
		Context: ctx,
	}
}

// NewPostLibrarianUsernameBookParamsWithHTTPClient creates a new PostLibrarianUsernameBookParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLibrarianUsernameBookParamsWithHTTPClient(client *http.Client) *PostLibrarianUsernameBookParams {
	return &PostLibrarianUsernameBookParams{
		HTTPClient: client,
	}
}

/* PostLibrarianUsernameBookParams contains all the parameters to send to the API endpoint
   for the post librarian username book operation.

   Typically these are written to a http.Request.
*/
type PostLibrarianUsernameBookParams struct {

	// Book.
	Book *models.Book

	// Username.
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post librarian username book params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLibrarianUsernameBookParams) WithDefaults() *PostLibrarianUsernameBookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post librarian username book params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLibrarianUsernameBookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) WithTimeout(timeout time.Duration) *PostLibrarianUsernameBookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) WithContext(ctx context.Context) *PostLibrarianUsernameBookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) WithHTTPClient(client *http.Client) *PostLibrarianUsernameBookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBook adds the book to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) WithBook(book *models.Book) *PostLibrarianUsernameBookParams {
	o.SetBook(book)
	return o
}

// SetBook adds the book to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) SetBook(book *models.Book) {
	o.Book = book
}

// WithUsername adds the username to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) WithUsername(username string) *PostLibrarianUsernameBookParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the post librarian username book params
func (o *PostLibrarianUsernameBookParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PostLibrarianUsernameBookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Book != nil {
		if err := r.SetBodyParam(o.Book); err != nil {
			return err
		}
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
