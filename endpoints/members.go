package endpoints

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/mvo5/qrest-skeleton/backend"

	"github.com/MicahParks/qrest/models"
	"github.com/MicahParks/qrest/restapi/operations/api"
)

// HandleMembersAdd TODO
func HandleMembersAdd(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupMembersAddHandlerFunc {
	return func(params api.GroupMembersAddParams) middleware.Responder {

		// Iterate through the map of quota-groups to members.
		var err error
		for groupName, members := range params.GroupMembers {

			// Get the group from the quota manager.
			var group *backend.QuotaGroup
			if group = quotaManager.GetGroup(groupName); group == nil {
				// TODO
			}

			// Iterate through the member quota-groups.
			for _, memberGroup := range members.SubGroups {

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

// HandleMemberDelete TODO
func HandleMembersDelete(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupMembersDeleteHandlerFunc {
	return func(params api.GroupMembersDeleteParams) middleware.Responder {

		// Iterate through the map of quota-groups to members.
		var err error
		for groupName, members := range params.GroupMembers {

			// Get the group from the quota manager.
			var group *backend.QuotaGroup
			if group = quotaManager.GetGroup(groupName); group == nil {
				// TODO
			}

			// Iterate through the member quota-groups.
			for _, memberGroup := range members.SubGroups {

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

// memberAddRequestFailure TODO
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

	// Create the response back to the client.
	code := 500
	errPayload := &models.Error{
		Code:    int64(code),
		Message: message,
	}

	// Decide which Go type to return.
	if !addMember {
		resp := &api.GroupMembersDeleteDefault{
			Payload: errPayload,
		}
		resp.SetStatusCode(code)
		return resp
	}
	resp := &api.GroupMembersAddDefault{
		Payload: errPayload,
	}
	resp.SetStatusCode(code)
	return resp
}
