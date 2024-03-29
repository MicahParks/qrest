package endpoints

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/mvo5/qrest-skeleton/backend"

	"github.com/MicahParks/qrest/models"
	"github.com/MicahParks/qrest/restapi/operations/api"
)

// HandleMembersAdd creates an endpoint handler via a closure that will add members to quota-groups when used.
func HandleMembersAdd(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupMembersAddHandlerFunc {
	return func(params api.GroupMembersAddParams) middleware.Responder {

		// Debug info.
		logger.Debugw("Touched.",
			"groups", params.GroupMembers,
		)

		// Iterate through the map of quota-groups to members.
		var err error
		for groupName, members := range params.GroupMembers {

			// Get the group from the quota manager.
			var group *backend.QuotaGroup
			if group = quotaManager.GetGroup(groupName); group == nil {

				// Log the event.
				code, message := groupNotFound(groupName)
				logger.Infow(message,
					"groupName", groupName,
				)

				// Report the error to the client.
				return errorResponse(code, message, &api.GroupMembersAddDefault{})
			}

			// Iterate through the member quota-groups.
			for _, memberGroup := range members.MemberGroups {

				// Add the member quota-group to the quota-group.
				if err = group.AddGroup(memberGroup); err != nil {
					return memberFailure(true, err, groupName, logger, memberGroup, "quota-group")
				}
			}

			// Iterate through the member snaps.
			for _, snap := range members.Snaps {

				// Add the member snap to the quota-group.
				if err = group.AddGroup(snap); err != nil {
					return memberFailure(true, err, groupName, logger, snap, "snap")
				}
			}
		}

		return &api.GroupMembersAddOK{}
	}
}

// HandleMemberDelete creates an endpoint handler via a closure that will delete members from quota-groups when used.
func HandleMembersDelete(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupMembersDeleteHandlerFunc {
	return func(params api.GroupMembersDeleteParams) middleware.Responder {

		// Debug info.
		logger.Debugw("Touched.",
			"groupMembers", params.GroupMembers,
		)

		// Iterate through the map of quota-groups to members.
		var err error
		for groupName, members := range params.GroupMembers {

			// Get the group from the quota manager.
			var group *backend.QuotaGroup
			if group = quotaManager.GetGroup(groupName); group == nil {

				// Log the event.
				code, message := groupNotFound(groupName)
				logger.Infow(message,
					"groupName", groupName,
				)

				// Report the error to the client.
				return errorResponse(code, message, &api.GroupMembersDeleteDefault{})
			}

			// Iterate through the member quota-groups.
			for _, memberGroup := range members.MemberGroups {

				// Add the member quota-group to the quota-group.
				if err = group.RemoveGroup(memberGroup); err != nil {
					return memberFailure(false, err, groupName, logger, memberGroup, "quota-group")
				}
			}

			// Iterate through the member snaps.
			for _, snap := range members.Snaps {

				// Add the member snap to the quota-group.
				if err = group.RemoveSnap(snap); err != nil {
					return memberFailure(false, err, groupName, logger, snap, "snap")
				}
			}
		}

		return &api.GroupMembersDeleteOK{}
	}
}

// HandleMemberRead creates an endpoint handler via a closure that will read members from quota-groups when used. This
// is an extra endpoint that was not asked for.
func HandleMembersRead(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupMembersReadHandlerFunc {
	return func(params api.GroupMembersReadParams) middleware.Responder {

		// Debug info.
		logger.Debugw("Touched.",
			"groups", params.Groups,
		)

		// Create the map returned to the client.
		groupMembers := make(map[string]models.GroupMembers)

		// Iterate through the given groups.
		for _, groupName := range params.Groups {

			// Get the group from the quota manager.
			var group *backend.QuotaGroup
			if group = quotaManager.GetGroup(groupName); group == nil {

				// Log the event.
				code, message := groupNotFound(groupName)
				logger.Infow(message,
					"groupName", groupName,
				)

				// Report the error to the client.
				return errorResponse(code, message, &api.GroupMembersReadDefault{})
			}

			// Add the group's group members to the map.
			groupMembers[groupName] = models.GroupMembers{
				Snaps:        group.Snaps(),
				MemberGroups: group.Groups(),
			}
		}

		return &api.GroupMembersReadOK{
			Payload: groupMembers,
		}
	}
}

// memberAddRequestFailure helps cut down on duplicate code by logging an error and creating the appropriate response
// type.
func memberFailure(addMember bool, err error, groupName string, logger *zap.SugaredLogger, memberName, memberType string) middleware.Responder {

	// Log the error.
	operation := "add"
	if !addMember {
		operation = "delete"
	}
	message := fmt.Sprintf("Failed to %s member \"%s\".", operation, memberName)
	logger.Infow(message,
		"quota-group", groupName,
		"memberName", memberName,
		"memberType", memberType,
		"error", err.Error(),
	)

	// Decide which Go type to return and create the response back to the client..
	code := 500
	if !addMember {
		return errorResponse(code, message, &api.GroupMembersDeleteDefault{})
	}
	return errorResponse(code, message, &api.GroupMembersAddDefault{})
}
