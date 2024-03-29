// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/MicahParks/qrest/configure"
	"github.com/MicahParks/qrest/endpoints"
	"github.com/MicahParks/qrest/endpoints/system"
	"github.com/MicahParks/qrest/restapi/operations"
)

//go:generate swagger generate server --target ../../qrest-skeleton --name SnapAPIChallenge --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.SnapAPIChallengeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SnapAPIChallengeAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Configure the service.
	quotaManager, logger, err := configure.Configure()
	if err != nil {
		log.Fatalf("Failed to configure the service.\nError: %s", err.Error())
	}

	// Set the generated code logger.
	api.Logger = logger.Named("Generated Code").Infof

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Set the endpoint handlers.
	api.SystemAliveHandler = system.HandleAlive()
	api.APIGroupDeleteHandler = endpoints.HandleGroupDelete(logger.Named("DELETE /group"), quotaManager)
	api.APIGroupInsertHandler = endpoints.HandleGroupInsert(logger.Named("POST /group"), quotaManager)
	api.APIGroupLimitReadHandler = endpoints.HandleGroupLimitsRead(logger.Named("GET /group/limits"), quotaManager)
	api.APIGroupLimitWriteHandler = endpoints.HandleGroupLimitsWrite(logger.Named("POST /group/limits"), quotaManager)
	api.APIGroupMembersAddHandler = endpoints.HandleMembersAdd(logger.Named("POST /group/members"), quotaManager)
	api.APIGroupMembersDeleteHandler = endpoints.HandleMembersDelete(logger.Named("DELETE /group/members"), quotaManager)
	api.APIGroupMembersReadHandler = endpoints.HandleMembersRead(logger.Named("GET /group/members"), quotaManager)
	api.APIGroupUsageHandler = endpoints.HandleUsage(logger.Named("GET /group/usage"), quotaManager)

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

	// Create an incoming request rate limiter that only allows 1 request per section and forgets about clients after 1
	// hour.
	limit := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	// Find the IP of the client in the X-Forwarded-For header, because Caddy will be the server in front of this.
	limit.SetIPLookups([]string{"X-Forwarded-For"})

	// Follow the HTTP middleware pattern.
	return tollbooth.LimitHandler(limit, handler) // TODO Logging middleware. Maybe another rate limiter instead.
}
