package endpoints

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"

	"go.uber.org/zap"

	"github.com/MicahParks/qrest/backend"
	"github.com/MicahParks/qrest/models"
	"github.com/MicahParks/qrest/restapi/operations/api"
)

// HandleGroupDelete TODO
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
