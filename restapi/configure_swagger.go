// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/mvo5/qrest-skeleton/restapi/operations"
	"github.com/mvo5/qrest-skeleton/restapi/operations/system"
)

//go:generate swagger generate server --target ../../qrest-skeleton --name Swagger --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.SwaggerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SwaggerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.SystemAliveHandler == nil {
		api.SystemAliveHandler = system.AliveHandlerFunc(func(params system.AliveParams) middleware.Responder {
			return middleware.NotImplemented("operation system.Alive has not yet been implemented")
		})
	}
	if api.GroupDeleteHandler == nil {
		api.GroupDeleteHandler = operations.GroupDeleteHandlerFunc(func(params operations.GroupDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GroupDelete has not yet been implemented")
		})
	}
	if api.GroupInsertHandler == nil {
		api.GroupInsertHandler = operations.GroupInsertHandlerFunc(func(params operations.GroupInsertParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GroupInsert has not yet been implemented")
		})
	}
	if api.GroupLimitReadHandler == nil {
		api.GroupLimitReadHandler = operations.GroupLimitReadHandlerFunc(func(params operations.GroupLimitReadParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GroupLimitRead has not yet been implemented")
		})
	}
	if api.GroupLimitWriteHandler == nil {
		api.GroupLimitWriteHandler = operations.GroupLimitWriteHandlerFunc(func(params operations.GroupLimitWriteParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GroupLimitWrite has not yet been implemented")
		})
	}
	if api.GroupMembersAddHandler == nil {
		api.GroupMembersAddHandler = operations.GroupMembersAddHandlerFunc(func(params operations.GroupMembersAddParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GroupMembersAdd has not yet been implemented")
		})
	}
	if api.GroupMembersDeleteHandler == nil {
		api.GroupMembersDeleteHandler = operations.GroupMembersDeleteHandlerFunc(func(params operations.GroupMembersDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GroupMembersDelete has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
