package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mvo5/qrest-skeleton/backend"
	"go.uber.org/zap"

	"github.com/MicahParks/qrest/models"
	"github.com/MicahParks/qrest/restapi/operations/api"
)

// HandleGroupLimitsRead TODO
func HandleGroupLimitsRead(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupLimitReadHandlerFunc {
	return func(params api.GroupLimitReadParams) middleware.Responder {

		// Create the map returned to the client.
		groupLimits := make(map[string]models.Limits)

		// Iterate through the given quota-groups.
		for _, groupName := range params.Group {

			// Check to see if the group is a duplicate.
			if _, ok := groupLimits[groupName]; ok {
				continue
			}

			// Get the group from the quota manager.
			var group *backend.QuotaGroup
			if group = quotaManager.GetGroup(groupName); group == nil {
				// TODO
			}

			// Create the limits Go structure.
			limits := models.Limits{
				MaxMemory: group.MaxMemory(),
			}

			// Add the group and its limits to the map.
			groupLimits[groupName] = limits
		}

		return &api.GroupLimitReadOK{
			Payload: groupLimits,
		}
	}
}

// HandleGroupLimitsWrite TODO
func HandleGroupLimitsWrite(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupLimitWriteHandlerFunc {
	return func(params api.GroupLimitWriteParams) middleware.Responder {

		// Iterate through the given quota-groups and their limits.
		var err error
		for groupName, limits := range params.GroupLimitsMap {

			// Get the group from the quota manager.
			var group *backend.QuotaGroup
			if group = quotaManager.GetGroup(groupName); group == nil {
				// TODO
			}

			// Set the maximum memory for the quota-group.
			if err = group.SetMaxMemory(limits.MaxMemory); err != nil {
				// TODO
			}
		}

		return &api.GroupLimitWriteOK{}
	}
}
