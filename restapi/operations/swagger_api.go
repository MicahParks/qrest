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

	apiops "github.com/MicahParks/qrest/restapi/operations/api"
	"github.com/MicahParks/qrest/restapi/operations/system"
)

// NewSwaggerAPI creates a new Swagger instance
func NewSwaggerAPI(spec *loads.Document) *SwaggerAPI {
	return &SwaggerAPI{
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

		SystemAliveHandler: system.AliveHandlerFunc(func(params system.AliveParams) middleware.Responder {
			return middleware.NotImplemented("operation system.Alive has not yet been implemented")
		}),
		APIGroupDeleteHandler: apiops.GroupDeleteHandlerFunc(func(params apiops.GroupDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation api.GroupDelete has not yet been implemented")
		}),
		APIGroupInsertHandler: apiops.GroupInsertHandlerFunc(func(params apiops.GroupInsertParams) middleware.Responder {
			return middleware.NotImplemented("operation api.GroupInsert has not yet been implemented")
		}),
		APIGroupLimitReadHandler: apiops.GroupLimitReadHandlerFunc(func(params apiops.GroupLimitReadParams) middleware.Responder {
			return middleware.NotImplemented("operation api.GroupLimitRead has not yet been implemented")
		}),
		APIGroupLimitWriteHandler: apiops.GroupLimitWriteHandlerFunc(func(params apiops.GroupLimitWriteParams) middleware.Responder {
			return middleware.NotImplemented("operation api.GroupLimitWrite has not yet been implemented")
		}),
		APIGroupMembersAddHandler: apiops.GroupMembersAddHandlerFunc(func(params apiops.GroupMembersAddParams) middleware.Responder {
			return middleware.NotImplemented("operation api.GroupMembersAdd has not yet been implemented")
		}),
		APIGroupMembersDeleteHandler: apiops.GroupMembersDeleteHandlerFunc(func(params apiops.GroupMembersDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation api.GroupMembersDelete has not yet been implemented")
		}),
		APIGroupUsageHandler: apiops.GroupUsageHandlerFunc(func(params apiops.GroupUsageParams) middleware.Responder {
			return middleware.NotImplemented("operation api.GroupUsage has not yet been implemented")
		}),
	}
}

/*SwaggerAPI the swagger API */
type SwaggerAPI struct {
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

	// SystemAliveHandler sets the operation handler for the alive operation
	SystemAliveHandler system.AliveHandler
	// APIGroupDeleteHandler sets the operation handler for the group delete operation
	APIGroupDeleteHandler apiops.GroupDeleteHandler
	// APIGroupInsertHandler sets the operation handler for the group insert operation
	APIGroupInsertHandler apiops.GroupInsertHandler
	// APIGroupLimitReadHandler sets the operation handler for the group limit read operation
	APIGroupLimitReadHandler apiops.GroupLimitReadHandler
	// APIGroupLimitWriteHandler sets the operation handler for the group limit write operation
	APIGroupLimitWriteHandler apiops.GroupLimitWriteHandler
	// APIGroupMembersAddHandler sets the operation handler for the group members add operation
	APIGroupMembersAddHandler apiops.GroupMembersAddHandler
	// APIGroupMembersDeleteHandler sets the operation handler for the group members delete operation
	APIGroupMembersDeleteHandler apiops.GroupMembersDeleteHandler
	// APIGroupUsageHandler sets the operation handler for the group usage operation
	APIGroupUsageHandler apiops.GroupUsageHandler

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
func (o *SwaggerAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *SwaggerAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *SwaggerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SwaggerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SwaggerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SwaggerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SwaggerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SwaggerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SwaggerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SwaggerAPI
func (o *SwaggerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.SystemAliveHandler == nil {
		unregistered = append(unregistered, "system.AliveHandler")
	}
	if o.APIGroupDeleteHandler == nil {
		unregistered = append(unregistered, "api.GroupDeleteHandler")
	}
	if o.APIGroupInsertHandler == nil {
		unregistered = append(unregistered, "api.GroupInsertHandler")
	}
	if o.APIGroupLimitReadHandler == nil {
		unregistered = append(unregistered, "api.GroupLimitReadHandler")
	}
	if o.APIGroupLimitWriteHandler == nil {
		unregistered = append(unregistered, "api.GroupLimitWriteHandler")
	}
	if o.APIGroupMembersAddHandler == nil {
		unregistered = append(unregistered, "api.GroupMembersAddHandler")
	}
	if o.APIGroupMembersDeleteHandler == nil {
		unregistered = append(unregistered, "api.GroupMembersDeleteHandler")
	}
	if o.APIGroupUsageHandler == nil {
		unregistered = append(unregistered, "api.GroupUsageHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SwaggerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SwaggerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *SwaggerAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *SwaggerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *SwaggerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *SwaggerAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the swagger API
func (o *SwaggerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SwaggerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alive"] = system.NewAlive(o.context, o.SystemAliveHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/group"] = apiops.NewGroupDelete(o.context, o.APIGroupDeleteHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/group"] = apiops.NewGroupInsert(o.context, o.APIGroupInsertHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/group/limits"] = apiops.NewGroupLimitRead(o.context, o.APIGroupLimitReadHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/group/limits"] = apiops.NewGroupLimitWrite(o.context, o.APIGroupLimitWriteHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/group/members"] = apiops.NewGroupMembersAdd(o.context, o.APIGroupMembersAddHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/group/members"] = apiops.NewGroupMembersDelete(o.context, o.APIGroupMembersDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/group/usage"] = apiops.NewGroupUsage(o.context, o.APIGroupUsageHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SwaggerAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *SwaggerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SwaggerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SwaggerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *SwaggerAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
