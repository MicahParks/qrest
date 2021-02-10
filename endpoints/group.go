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
func HandleGroupInsert(_ *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupInsertHandlerFunc {
	return func(params api.GroupInsertParams) middleware.Responder {

		// Iterate through the given groups.
		for _, group := range params.Group {
			if group != nil {

				// Validate the input (more than the default amount).
				if group.Name == "" || group.Limits == nil {

					// Report the group as unprocessable.
					code := 422
					message := fmt.Sprintf("The group \"%s\" was not processable. Empty group names or nil resource limits are not allowed.", group.Name)
					resp := &api.GroupInsertDefault{Payload: &models.Error{
						Code:    int64(code),
						Message: message,
					}}
					resp.SetStatusCode(code)

					return resp
				}

				// Add the group to the quota manager.
				quotaManager.AddGroup(group.Name, group.Limits.MaxMemory)
			}
		}

		return &api.GroupInsertOK{}
	}
}
