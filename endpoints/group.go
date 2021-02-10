package endpoints

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"

	"go.uber.org/zap"

	"github.com/mvo5/qrest-skeleton/backend"

	"github.com/MicahParks/qrest/restapi/operations/api"
)

// HandleGroupDelete creates an endpoint handler via a closure that will delete quota-groups when used.
func HandleGroupDelete(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupDeleteHandlerFunc {
	return func(params api.GroupDeleteParams) middleware.Responder {

		// Debug info.
		logger.Debugw("Touched.",
			"groups", params.Groups,
		)

		// Iterate through the given groups.
		var err error
		for _, name := range params.Groups {

			// Remove the group from the quota manager.
			if err = quotaManager.RemoveGroup(name); err != nil {

				// Log the error.
				message := fmt.Sprintf("Failed to remove group \"%s\" from quota manager.", name)
				logger.Infow(message,
					"name", name,
					"error", err.Error(),
				)

				// Report the error to the client.
				return errorResponse(500, message, &api.GroupDeleteDefault{})
			}
		}

		return &api.GroupDeleteOK{}
	}
}

// HandleGroupInsert creates an endpoint handler via a closure that will insert quota-groups when used.
func HandleGroupInsert(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupInsertHandlerFunc {
	return func(params api.GroupInsertParams) middleware.Responder {

		// Debug info.
		logger.Debugw("Touched.",
			"groups", params.Groups,
		)

		// Iterate through the given groups.
		var err error
		for _, inputGroup := range params.Groups {
			if inputGroup != nil {

				// Validate the input (more than the default amount).
				if inputGroup.Name == "" {

					// Report the group as unprocessable.
					message := fmt.Sprintf("The group \"%s\" was not processable. Empty group names are not allowed.", inputGroup.Name)
					return errorResponse(422, message, &api.GroupInsertDefault{})
				}

				// If no limit was provided, use 0.
				var memoryLimit uint64
				if inputGroup.Limits != nil {
					memoryLimit = inputGroup.Limits.MaxMemory
				}

				// Add the group to the quota manager.
				group := quotaManager.AddGroup(inputGroup.Name, memoryLimit)

				// Iterate through the input quota-group's member quota-groups.
				for _, groupName := range inputGroup.SubGroups {

					// Confirm the group exists in the quoteManager.
					if quotaManager.GetGroup(groupName) != nil {
						if err = group.AddGroup(groupName); err != nil {

							// Log the event.
							message := fmt.Sprintf("Could not add \"%s\" member quota-group.", groupName)
							logger.Infow(message,
								"groupName", groupName,
								"error", err.Error(),
							)

							// Report the error to the client.
							return errorResponse(500, message, &api.GroupInsertDefault{})
						}
					} else {

						// Report the error to the client.
						message := fmt.Sprintf("Could not add \"%s\" member quota-group because it is not tracked by the quota manager.", groupName)
						return errorResponse(400, message, &api.GroupInsertDefault{})
					}
				}

				// Iterate through the input quota-group's snaps.
				for _, snap := range inputGroup.Snaps {

					// Add the snap to the quota-group.
					group.AddSnap(snap)
				}
			}
		}

		return &api.GroupInsertOK{}
	}
}
