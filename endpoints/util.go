package endpoints

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/MicahParks/qrest/models"
)

// defaultResponse is an interface used to pass different types of default responses and return an error responder.
type defaultResponse interface {
	SetStatusCode(code int)
	SetPayload(payload *models.Error)
	WriteResponse(rw http.ResponseWriter, producer runtime.Producer)
}

// errorResponse creates a response given the required assets.
func errorResponse(code int, message string, resp defaultResponse) middleware.Responder {

	// Set the payload for the response.
	resp.SetPayload(&models.Error{
		Code:    int64(code),
		Message: message,
	})

	// Set the status code for the response.
	resp.SetStatusCode(code)

	return resp
}

// groupNotFound creates an HTTP status code and message for when a group is not found.
func groupNotFound(groupName string) (code int, message string) {
	return 400, fmt.Sprintf("The quota-group \"%s\" was not found.", groupName)
}
