package endpoints

import (
	"math/rand"

	"github.com/go-openapi/runtime/middleware"
	"github.com/mvo5/qrest-skeleton/backend"
	"go.uber.org/zap"

	"github.com/MicahParks/qrest/models"
	"github.com/MicahParks/qrest/restapi/operations/api"
)

const (

	// memoryUsageKey represents the key in a resource usage map that stores memory usage.
	memoryUsageKey = "memory"
)

// HandleUsage creates an endpoint handler via a closure that will read quota-group usage when used.
func HandleUsage(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupUsageHandlerFunc {
	return func(params api.GroupUsageParams) middleware.Responder {

		// Debug info.
		logger.Debugw("Touched.",
			"groups", params.Groups,
		)

		// Create the map to return to the client.
		groupUsage := make(map[string]models.Usage)

		// Iterate through the given groups.
		for _, groupName := range params.Groups {

			// Create the usage map for this group.
			usage := make(map[string]models.UsageAnon)

			// Get the group from the quota manager.
			var group *backend.QuotaGroup
			if group = quotaManager.GetGroup(groupName); group == nil {

				// Log the event.
				code, message := groupNotFound(groupName)
				logger.Infow(message,
					"groupName", groupName,
				)

				// Report the error to the client.
				return errorResponse(code, message, &api.GroupUsageDefault{})
			}

			// Get the memory utilization.
			max := float64(group.MaxMemory()) // TODO Check for data loss due to conversion.
			usage[memoryUsageKey] = models.UsageAnon{
				Max:   max,
				Usage: max * rand.Float64(), // TODO Implement this on the backend.
			}

			// Add the group's usage to the map.
			groupUsage[groupName] = usage
		}

		return &api.GroupUsageOK{
			Payload: groupUsage,
		}
	}
}
