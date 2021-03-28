// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetLibrarianUsernameBookTitleAuthorHandlerFunc turns a function with the right signature into a get librarian username book title author handler
type GetLibrarianUsernameBookTitleAuthorHandlerFunc func(GetLibrarianUsernameBookTitleAuthorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLibrarianUsernameBookTitleAuthorHandlerFunc) Handle(params GetLibrarianUsernameBookTitleAuthorParams) middleware.Responder {
	return fn(params)
}

// GetLibrarianUsernameBookTitleAuthorHandler interface for that can handle valid get librarian username book title author params
type GetLibrarianUsernameBookTitleAuthorHandler interface {
	Handle(GetLibrarianUsernameBookTitleAuthorParams) middleware.Responder
}

// NewGetLibrarianUsernameBookTitleAuthor creates a new http.Handler for the get librarian username book title author operation
func NewGetLibrarianUsernameBookTitleAuthor(ctx *middleware.Context, handler GetLibrarianUsernameBookTitleAuthorHandler) *GetLibrarianUsernameBookTitleAuthor {
	return &GetLibrarianUsernameBookTitleAuthor{Context: ctx, Handler: handler}
}

/* GetLibrarianUsernameBookTitleAuthor swagger:route GET /librarian/{username}/book/{title}/{author} getLibrarianUsernameBookTitleAuthor

request a books stock information by its title and author

*/
type GetLibrarianUsernameBookTitleAuthor struct {
	Context *middleware.Context
	Handler GetLibrarianUsernameBookTitleAuthorHandler
}

func (o *GetLibrarianUsernameBookTitleAuthor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLibrarianUsernameBookTitleAuthorParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
