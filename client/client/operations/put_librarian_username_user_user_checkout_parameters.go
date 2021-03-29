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

// NewPutLibrarianUsernameUserUserCheckoutParams creates a new PutLibrarianUsernameUserUserCheckoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutLibrarianUsernameUserUserCheckoutParams() *PutLibrarianUsernameUserUserCheckoutParams {
	return &PutLibrarianUsernameUserUserCheckoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutLibrarianUsernameUserUserCheckoutParamsWithTimeout creates a new PutLibrarianUsernameUserUserCheckoutParams object
// with the ability to set a timeout on a request.
func NewPutLibrarianUsernameUserUserCheckoutParamsWithTimeout(timeout time.Duration) *PutLibrarianUsernameUserUserCheckoutParams {
	return &PutLibrarianUsernameUserUserCheckoutParams{
		timeout: timeout,
	}
}

// NewPutLibrarianUsernameUserUserCheckoutParamsWithContext creates a new PutLibrarianUsernameUserUserCheckoutParams object
// with the ability to set a context for a request.
func NewPutLibrarianUsernameUserUserCheckoutParamsWithContext(ctx context.Context) *PutLibrarianUsernameUserUserCheckoutParams {
	return &PutLibrarianUsernameUserUserCheckoutParams{
		Context: ctx,
	}
}

// NewPutLibrarianUsernameUserUserCheckoutParamsWithHTTPClient creates a new PutLibrarianUsernameUserUserCheckoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutLibrarianUsernameUserUserCheckoutParamsWithHTTPClient(client *http.Client) *PutLibrarianUsernameUserUserCheckoutParams {
	return &PutLibrarianUsernameUserUserCheckoutParams{
		HTTPClient: client,
	}
}

/* PutLibrarianUsernameUserUserCheckoutParams contains all the parameters to send to the API endpoint
   for the put librarian username user user checkout operation.

   Typically these are written to a http.Request.
*/
type PutLibrarianUsernameUserUserCheckoutParams struct {

	// Book.
	Book PutLibrarianUsernameUserUserCheckoutBody

	// User.
	User string

	// Username.
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put librarian username user user checkout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLibrarianUsernameUserUserCheckoutParams) WithDefaults() *PutLibrarianUsernameUserUserCheckoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put librarian username user user checkout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLibrarianUsernameUserUserCheckoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) WithTimeout(timeout time.Duration) *PutLibrarianUsernameUserUserCheckoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) WithContext(ctx context.Context) *PutLibrarianUsernameUserUserCheckoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) WithHTTPClient(client *http.Client) *PutLibrarianUsernameUserUserCheckoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBook adds the book to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) WithBook(book PutLibrarianUsernameUserUserCheckoutBody) *PutLibrarianUsernameUserUserCheckoutParams {
	o.SetBook(book)
	return o
}

// SetBook adds the book to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) SetBook(book PutLibrarianUsernameUserUserCheckoutBody) {
	o.Book = book
}

// WithUser adds the user to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) WithUser(user string) *PutLibrarianUsernameUserUserCheckoutParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) SetUser(user string) {
	o.User = user
}

// WithUsername adds the username to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) WithUsername(username string) *PutLibrarianUsernameUserUserCheckoutParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the put librarian username user user checkout params
func (o *PutLibrarianUsernameUserUserCheckoutParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PutLibrarianUsernameUserUserCheckoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Book); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
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