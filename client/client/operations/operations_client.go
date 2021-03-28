// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteBookTitle(params *DeleteBookTitleParams, opts ...ClientOption) (*DeleteBookTitleOK, error)

	DeleteLibrarianUsernameBookTitleAuthorCopies(params *DeleteLibrarianUsernameBookTitleAuthorCopiesParams, opts ...ClientOption) (*DeleteLibrarianUsernameBookTitleAuthorCopiesOK, error)

	DeleteLibrarianUsernameUserUser(params *DeleteLibrarianUsernameUserUserParams, opts ...ClientOption) (*DeleteLibrarianUsernameUserUserOK, error)

	GetBook(params *GetBookParams, opts ...ClientOption) (*GetBookOK, error)

	GetBookTitle(params *GetBookTitleParams, opts ...ClientOption) (*GetBookTitleOK, error)

	GetLibrarianUsernameBookTitleAuthor(params *GetLibrarianUsernameBookTitleAuthorParams, opts ...ClientOption) (*GetLibrarianUsernameBookTitleAuthorOK, error)

	GetLibrarianUsernameUserUser(params *GetLibrarianUsernameUserUserParams, opts ...ClientOption) (*GetLibrarianUsernameUserUserOK, error)

	PostBook(params *PostBookParams, opts ...ClientOption) (*PostBookCreated, error)

	PostLibrarianUsernameBook(params *PostLibrarianUsernameBookParams, opts ...ClientOption) (*PostLibrarianUsernameBookCreated, error)

	PostLibrarianUsernameUser(params *PostLibrarianUsernameUserParams, opts ...ClientOption) (*PostLibrarianUsernameUserCreated, error)

	PutBookTitle(params *PutBookTitleParams, opts ...ClientOption) (*PutBookTitleOK, error)

	PutLibrarianUsernameUserUserCheckout(params *PutLibrarianUsernameUserUserCheckoutParams, opts ...ClientOption) (*PutLibrarianUsernameUserUserCheckoutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteBookTitle deletes a book by its title
*/
func (a *Client) DeleteBookTitle(params *DeleteBookTitleParams, opts ...ClientOption) (*DeleteBookTitleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBookTitleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBookTitle",
		Method:             "DELETE",
		PathPattern:        "/book/{title}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBookTitleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBookTitleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteBookTitle: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteLibrarianUsernameBookTitleAuthorCopies deletes a book by its title author
*/
func (a *Client) DeleteLibrarianUsernameBookTitleAuthorCopies(params *DeleteLibrarianUsernameBookTitleAuthorCopiesParams, opts ...ClientOption) (*DeleteLibrarianUsernameBookTitleAuthorCopiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLibrarianUsernameBookTitleAuthorCopiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteLibrarianUsernameBookTitleAuthorCopies",
		Method:             "DELETE",
		PathPattern:        "/librarian/{username}/book/{title}/{author}/{copies}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLibrarianUsernameBookTitleAuthorCopiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLibrarianUsernameBookTitleAuthorCopiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteLibrarianUsernameBookTitleAuthorCopies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteLibrarianUsernameUserUser deletes a user by their username
*/
func (a *Client) DeleteLibrarianUsernameUserUser(params *DeleteLibrarianUsernameUserUserParams, opts ...ClientOption) (*DeleteLibrarianUsernameUserUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLibrarianUsernameUserUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteLibrarianUsernameUserUser",
		Method:             "DELETE",
		PathPattern:        "/librarian/{username}/user/{user}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLibrarianUsernameUserUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLibrarianUsernameUserUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteLibrarianUsernameUserUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBook requests a book by its title
*/
func (a *Client) GetBook(params *GetBookParams, opts ...ClientOption) (*GetBookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBook",
		Method:             "GET",
		PathPattern:        "/book",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBookReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBook: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBookTitle requests a book by its title
*/
func (a *Client) GetBookTitle(params *GetBookTitleParams, opts ...ClientOption) (*GetBookTitleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBookTitleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBookTitle",
		Method:             "GET",
		PathPattern:        "/book/{title}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBookTitleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBookTitleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBookTitle: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLibrarianUsernameBookTitleAuthor requests a books stock information by its title and author
*/
func (a *Client) GetLibrarianUsernameBookTitleAuthor(params *GetLibrarianUsernameBookTitleAuthorParams, opts ...ClientOption) (*GetLibrarianUsernameBookTitleAuthorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLibrarianUsernameBookTitleAuthorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLibrarianUsernameBookTitleAuthor",
		Method:             "GET",
		PathPattern:        "/librarian/{username}/book/{title}/{author}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLibrarianUsernameBookTitleAuthorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLibrarianUsernameBookTitleAuthorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLibrarianUsernameBookTitleAuthor: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLibrarianUsernameUserUser requests a user by their username
*/
func (a *Client) GetLibrarianUsernameUserUser(params *GetLibrarianUsernameUserUserParams, opts ...ClientOption) (*GetLibrarianUsernameUserUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLibrarianUsernameUserUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLibrarianUsernameUserUser",
		Method:             "GET",
		PathPattern:        "/librarian/{username}/user/{user}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLibrarianUsernameUserUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLibrarianUsernameUserUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLibrarianUsernameUserUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostBook requests a book be added to the library
*/
func (a *Client) PostBook(params *PostBookParams, opts ...ClientOption) (*PostBookCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostBook",
		Method:             "POST",
		PathPattern:        "/book",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBookReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostBookCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostBook: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostLibrarianUsernameBook requests a book be added to the library
*/
func (a *Client) PostLibrarianUsernameBook(params *PostLibrarianUsernameBookParams, opts ...ClientOption) (*PostLibrarianUsernameBookCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLibrarianUsernameBookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostLibrarianUsernameBook",
		Method:             "POST",
		PathPattern:        "/librarian/{username}/book",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLibrarianUsernameBookReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLibrarianUsernameBookCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLibrarianUsernameBook: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostLibrarianUsernameUser requests a user be added to the library
*/
func (a *Client) PostLibrarianUsernameUser(params *PostLibrarianUsernameUserParams, opts ...ClientOption) (*PostLibrarianUsernameUserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLibrarianUsernameUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostLibrarianUsernameUser",
		Method:             "POST",
		PathPattern:        "/librarian/{username}/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLibrarianUsernameUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLibrarianUsernameUserCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLibrarianUsernameUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutBookTitle requests a book be updated
*/
func (a *Client) PutBookTitle(params *PutBookTitleParams, opts ...ClientOption) (*PutBookTitleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBookTitleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutBookTitle",
		Method:             "PUT",
		PathPattern:        "/book/{title}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBookTitleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutBookTitleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutBookTitle: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutLibrarianUsernameUserUserCheckout requests a book to be checked out by a user
*/
func (a *Client) PutLibrarianUsernameUserUserCheckout(params *PutLibrarianUsernameUserUserCheckoutParams, opts ...ClientOption) (*PutLibrarianUsernameUserUserCheckoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutLibrarianUsernameUserUserCheckoutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutLibrarianUsernameUserUserCheckout",
		Method:             "PUT",
		PathPattern:        "/librarian/{username}/user/{user}/checkout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutLibrarianUsernameUserUserCheckoutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutLibrarianUsernameUserUserCheckoutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutLibrarianUsernameUserUserCheckout: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
