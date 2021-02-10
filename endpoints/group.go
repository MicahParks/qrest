package endpoints

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"

	"go.uber.org/zap"

	"github.com/mvo5/qrest-skeleton/backend"

	"github.com/MicahParks/qrest/models"
	"github.com/MicahParks/qrest/restapi/operations/api"
)

// HandleGroupDelete creates an endpoint handler via a closure that will delete quota-groups when used.
func HandleGroupDelete(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupDeleteHandlerFunc {
	return func(params api.GroupDeleteParams) middleware.Responder {

		// Iterate through the given groups.
		var err error
		for _, name := range params.Group {

			// Remove the group from the quota manager.
			if err = quotaManager.RemoveGroup(name); err != nil {

				// Log the error.
				message := fmt.Sprintf("Failed to remove group \"%s\" from quota manager.", name)
				logger.Infow(message,
					"name", name,
					"error", err.Error(),
				)

				// Report the error to the client.
				code := 500
				resp := &api.GroupDeleteDefault{Payload: &models.Error{
					Code:    int64(code),
					Message: message,
				}}
				resp.SetStatusCode(code)

				return resp
			}
		}

		return &api.GroupDeleteOK{}
	}
}

// HandleGroupInsert creates an endpoint handler via a closure that will insert quota-groups when used.
func HandleGroupInsert(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupInsertHandlerFunc {
	return func(params api.GroupInsertParams) middleware.Responder {

		// Iterate through the given groups.
		for _, group := range params.Group {
			if group != nil {
				if group.Name == "" || group.Limits == nil {
					// TODO Return 422 Unprocessable.
				}

				// Add the group to the quota manager.
				quotaManager.AddGroup(group.Name, uint64(group.Limits.MaxMemory)) // TODO Integer conversion doesn't allow for full uint64.
			}
		}

		return &api.GroupInsertOK{}
	}
}
