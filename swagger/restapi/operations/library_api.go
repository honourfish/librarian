// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewLibraryAPI creates a new Library instance
func NewLibraryAPI(spec *loads.Document) *LibraryAPI {
	return &LibraryAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		DeleteBookTitleHandler: DeleteBookTitleHandlerFunc(func(params DeleteBookTitleParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteBookTitle has not yet been implemented")
		}),
		DeleteLibrarianUsernameBookTitleAuthorCopiesHandler: DeleteLibrarianUsernameBookTitleAuthorCopiesHandlerFunc(func(params DeleteLibrarianUsernameBookTitleAuthorCopiesParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteLibrarianUsernameBookTitleAuthorCopies has not yet been implemented")
		}),
		DeleteLibrarianUsernameUserUserHandler: DeleteLibrarianUsernameUserUserHandlerFunc(func(params DeleteLibrarianUsernameUserUserParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteLibrarianUsernameUserUser has not yet been implemented")
		}),
		GetBookHandler: GetBookHandlerFunc(func(params GetBookParams) middleware.Responder {
			return middleware.NotImplemented("operation GetBook has not yet been implemented")
		}),
		GetBookTitleHandler: GetBookTitleHandlerFunc(func(params GetBookTitleParams) middleware.Responder {
			return middleware.NotImplemented("operation GetBookTitle has not yet been implemented")
		}),
		GetLibrarianUsernameBookTitleAuthorHandler: GetLibrarianUsernameBookTitleAuthorHandlerFunc(func(params GetLibrarianUsernameBookTitleAuthorParams) middleware.Responder {
			return middleware.NotImplemented("operation GetLibrarianUsernameBookTitleAuthor has not yet been implemented")
		}),
		GetLibrarianUsernameUserUserHandler: GetLibrarianUsernameUserUserHandlerFunc(func(params GetLibrarianUsernameUserUserParams) middleware.Responder {
			return middleware.NotImplemented("operation GetLibrarianUsernameUserUser has not yet been implemented")
		}),
		PostBookHandler: PostBookHandlerFunc(func(params PostBookParams) middleware.Responder {
			return middleware.NotImplemented("operation PostBook has not yet been implemented")
		}),
		PostLibrarianUsernameBookHandler: PostLibrarianUsernameBookHandlerFunc(func(params PostLibrarianUsernameBookParams) middleware.Responder {
			return middleware.NotImplemented("operation PostLibrarianUsernameBook has not yet been implemented")
		}),
		PostLibrarianUsernameUserHandler: PostLibrarianUsernameUserHandlerFunc(func(params PostLibrarianUsernameUserParams) middleware.Responder {
			return middleware.NotImplemented("operation PostLibrarianUsernameUser has not yet been implemented")
		}),
		PutBookTitleHandler: PutBookTitleHandlerFunc(func(params PutBookTitleParams) middleware.Responder {
			return middleware.NotImplemented("operation PutBookTitle has not yet been implemented")
		}),
		PutLibrarianUsernameUserUserCheckoutHandler: PutLibrarianUsernameUserUserCheckoutHandlerFunc(func(params PutLibrarianUsernameUserUserCheckoutParams) middleware.Responder {
			return middleware.NotImplemented("operation PutLibrarianUsernameUserUserCheckout has not yet been implemented")
		}),
	}
}

/*LibraryAPI the library API */
type LibraryAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// DeleteBookTitleHandler sets the operation handler for the delete book title operation
	DeleteBookTitleHandler DeleteBookTitleHandler
	// DeleteLibrarianUsernameBookTitleAuthorCopiesHandler sets the operation handler for the delete librarian username book title author copies operation
	DeleteLibrarianUsernameBookTitleAuthorCopiesHandler DeleteLibrarianUsernameBookTitleAuthorCopiesHandler
	// DeleteLibrarianUsernameUserUserHandler sets the operation handler for the delete librarian username user user operation
	DeleteLibrarianUsernameUserUserHandler DeleteLibrarianUsernameUserUserHandler
	// GetBookHandler sets the operation handler for the get book operation
	GetBookHandler GetBookHandler
	// GetBookTitleHandler sets the operation handler for the get book title operation
	GetBookTitleHandler GetBookTitleHandler
	// GetLibrarianUsernameBookTitleAuthorHandler sets the operation handler for the get librarian username book title author operation
	GetLibrarianUsernameBookTitleAuthorHandler GetLibrarianUsernameBookTitleAuthorHandler
	// GetLibrarianUsernameUserUserHandler sets the operation handler for the get librarian username user user operation
	GetLibrarianUsernameUserUserHandler GetLibrarianUsernameUserUserHandler
	// PostBookHandler sets the operation handler for the post book operation
	PostBookHandler PostBookHandler
	// PostLibrarianUsernameBookHandler sets the operation handler for the post librarian username book operation
	PostLibrarianUsernameBookHandler PostLibrarianUsernameBookHandler
	// PostLibrarianUsernameUserHandler sets the operation handler for the post librarian username user operation
	PostLibrarianUsernameUserHandler PostLibrarianUsernameUserHandler
	// PutBookTitleHandler sets the operation handler for the put book title operation
	PutBookTitleHandler PutBookTitleHandler
	// PutLibrarianUsernameUserUserCheckoutHandler sets the operation handler for the put librarian username user user checkout operation
	PutLibrarianUsernameUserUserCheckoutHandler PutLibrarianUsernameUserUserCheckoutHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *LibraryAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *LibraryAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *LibraryAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *LibraryAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *LibraryAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *LibraryAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *LibraryAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *LibraryAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *LibraryAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the LibraryAPI
func (o *LibraryAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.DeleteBookTitleHandler == nil {
		unregistered = append(unregistered, "DeleteBookTitleHandler")
	}
	if o.DeleteLibrarianUsernameBookTitleAuthorCopiesHandler == nil {
		unregistered = append(unregistered, "DeleteLibrarianUsernameBookTitleAuthorCopiesHandler")
	}
	if o.DeleteLibrarianUsernameUserUserHandler == nil {
		unregistered = append(unregistered, "DeleteLibrarianUsernameUserUserHandler")
	}
	if o.GetBookHandler == nil {
		unregistered = append(unregistered, "GetBookHandler")
	}
	if o.GetBookTitleHandler == nil {
		unregistered = append(unregistered, "GetBookTitleHandler")
	}
	if o.GetLibrarianUsernameBookTitleAuthorHandler == nil {
		unregistered = append(unregistered, "GetLibrarianUsernameBookTitleAuthorHandler")
	}
	if o.GetLibrarianUsernameUserUserHandler == nil {
		unregistered = append(unregistered, "GetLibrarianUsernameUserUserHandler")
	}
	if o.PostBookHandler == nil {
		unregistered = append(unregistered, "PostBookHandler")
	}
	if o.PostLibrarianUsernameBookHandler == nil {
		unregistered = append(unregistered, "PostLibrarianUsernameBookHandler")
	}
	if o.PostLibrarianUsernameUserHandler == nil {
		unregistered = append(unregistered, "PostLibrarianUsernameUserHandler")
	}
	if o.PutBookTitleHandler == nil {
		unregistered = append(unregistered, "PutBookTitleHandler")
	}
	if o.PutLibrarianUsernameUserUserCheckoutHandler == nil {
		unregistered = append(unregistered, "PutLibrarianUsernameUserUserCheckoutHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *LibraryAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *LibraryAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *LibraryAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *LibraryAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *LibraryAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *LibraryAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the library API
func (o *LibraryAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *LibraryAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/book/{title}"] = NewDeleteBookTitle(o.context, o.DeleteBookTitleHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/librarian/{username}/book/{title}/{author}/{copies}"] = NewDeleteLibrarianUsernameBookTitleAuthorCopies(o.context, o.DeleteLibrarianUsernameBookTitleAuthorCopiesHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/librarian/{username}/user/{user}"] = NewDeleteLibrarianUsernameUserUser(o.context, o.DeleteLibrarianUsernameUserUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/book"] = NewGetBook(o.context, o.GetBookHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/book/{title}"] = NewGetBookTitle(o.context, o.GetBookTitleHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/librarian/{username}/book/{title}/{author}"] = NewGetLibrarianUsernameBookTitleAuthor(o.context, o.GetLibrarianUsernameBookTitleAuthorHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/librarian/{username}/user/{user}"] = NewGetLibrarianUsernameUserUser(o.context, o.GetLibrarianUsernameUserUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/book"] = NewPostBook(o.context, o.PostBookHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/librarian/{username}/book"] = NewPostLibrarianUsernameBook(o.context, o.PostLibrarianUsernameBookHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/librarian/{username}/user"] = NewPostLibrarianUsernameUser(o.context, o.PostLibrarianUsernameUserHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/book/{title}"] = NewPutBookTitle(o.context, o.PutBookTitleHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/librarian/{username}/user/{user}/checkout"] = NewPutLibrarianUsernameUserUserCheckout(o.context, o.PutLibrarianUsernameUserUserCheckoutHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *LibraryAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *LibraryAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *LibraryAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *LibraryAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *LibraryAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
