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

// NewGetLibrarianUsernameBookTitleAuthorParams creates a new GetLibrarianUsernameBookTitleAuthorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLibrarianUsernameBookTitleAuthorParams() *GetLibrarianUsernameBookTitleAuthorParams {
	return &GetLibrarianUsernameBookTitleAuthorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLibrarianUsernameBookTitleAuthorParamsWithTimeout creates a new GetLibrarianUsernameBookTitleAuthorParams object
// with the ability to set a timeout on a request.
func NewGetLibrarianUsernameBookTitleAuthorParamsWithTimeout(timeout time.Duration) *GetLibrarianUsernameBookTitleAuthorParams {
	return &GetLibrarianUsernameBookTitleAuthorParams{
		timeout: timeout,
	}
}

// NewGetLibrarianUsernameBookTitleAuthorParamsWithContext creates a new GetLibrarianUsernameBookTitleAuthorParams object
// with the ability to set a context for a request.
func NewGetLibrarianUsernameBookTitleAuthorParamsWithContext(ctx context.Context) *GetLibrarianUsernameBookTitleAuthorParams {
	return &GetLibrarianUsernameBookTitleAuthorParams{
		Context: ctx,
	}
}

// NewGetLibrarianUsernameBookTitleAuthorParamsWithHTTPClient creates a new GetLibrarianUsernameBookTitleAuthorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLibrarianUsernameBookTitleAuthorParamsWithHTTPClient(client *http.Client) *GetLibrarianUsernameBookTitleAuthorParams {
	return &GetLibrarianUsernameBookTitleAuthorParams{
		HTTPClient: client,
	}
}

/* GetLibrarianUsernameBookTitleAuthorParams contains all the parameters to send to the API endpoint
   for the get librarian username book title author operation.

   Typically these are written to a http.Request.
*/
type GetLibrarianUsernameBookTitleAuthorParams struct {

	// Author.
	Author string

	// Title.
	Title string

	// Username.
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get librarian username book title author params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLibrarianUsernameBookTitleAuthorParams) WithDefaults() *GetLibrarianUsernameBookTitleAuthorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get librarian username book title author params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLibrarianUsernameBookTitleAuthorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) WithTimeout(timeout time.Duration) *GetLibrarianUsernameBookTitleAuthorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) WithContext(ctx context.Context) *GetLibrarianUsernameBookTitleAuthorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) WithHTTPClient(client *http.Client) *GetLibrarianUsernameBookTitleAuthorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthor adds the author to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) WithAuthor(author string) *GetLibrarianUsernameBookTitleAuthorParams {
	o.SetAuthor(author)
	return o
}

// SetAuthor adds the author to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) SetAuthor(author string) {
	o.Author = author
}

// WithTitle adds the title to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) WithTitle(title string) *GetLibrarianUsernameBookTitleAuthorParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) SetTitle(title string) {
	o.Title = title
}

// WithUsername adds the username to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) WithUsername(username string) *GetLibrarianUsernameBookTitleAuthorParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get librarian username book title author params
func (o *GetLibrarianUsernameBookTitleAuthorParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetLibrarianUsernameBookTitleAuthorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param author
	if err := r.SetPathParam("author", o.Author); err != nil {
		return err
	}

	// path param title
	if err := r.SetPathParam("title", o.Title); err != nil {
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