// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutBookTitleHandlerFunc turns a function with the right signature into a put book title handler
type PutBookTitleHandlerFunc func(PutBookTitleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutBookTitleHandlerFunc) Handle(params PutBookTitleParams) middleware.Responder {
	return fn(params)
}

// PutBookTitleHandler interface for that can handle valid put book title params
type PutBookTitleHandler interface {
	Handle(PutBookTitleParams) middleware.Responder
}

// NewPutBookTitle creates a new http.Handler for the put book title operation
func NewPutBookTitle(ctx *middleware.Context, handler PutBookTitleHandler) *PutBookTitle {
	return &PutBookTitle{Context: ctx, Handler: handler}
}

/* PutBookTitle swagger:route PUT /book/{title} putBookTitle

request a book be updated

*/
type PutBookTitle struct {
	Context *middleware.Context
	Handler PutBookTitleHandler
}

func (o *PutBookTitle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutBookTitleParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}