// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBookTitleHandlerFunc turns a function with the right signature into a get book title handler
type GetBookTitleHandlerFunc func(GetBookTitleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBookTitleHandlerFunc) Handle(params GetBookTitleParams) middleware.Responder {
	return fn(params)
}

// GetBookTitleHandler interface for that can handle valid get book title params
type GetBookTitleHandler interface {
	Handle(GetBookTitleParams) middleware.Responder
}

// NewGetBookTitle creates a new http.Handler for the get book title operation
func NewGetBookTitle(ctx *middleware.Context, handler GetBookTitleHandler) *GetBookTitle {
	return &GetBookTitle{Context: ctx, Handler: handler}
}

/* GetBookTitle swagger:route GET /book/{title} getBookTitle

request a book by its title

*/
type GetBookTitle struct {
	Context *middleware.Context
	Handler GetBookTitleHandler
}

func (o *GetBookTitle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBookTitleParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
