package endpoints

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/qrest/backend"
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
			group := quotaManager.GetGroup(groupName)

			// Iterate through the member quota-groups.
			for _, memberGroup := range members.SubGroups {

				// Add the member quota-group to the quota-group.
				if err = group.AddGroup(memberGroup); err != nil {
					return memberRequestFailure(err, groupName, logger, memberGroup, "quota-group")
				}
			}

			// Iterate through the member snaps.
			for _, snap := range members.Snaps {

				// Add the member snap to the quota-group.
				if err = group.AddGroup(snap); err != nil {
					return memberRequestFailure(err, groupName, logger, snap, "snap")
				}
			}
		}

		return &api.GroupMembersAddOK{}
	}
}

// memberAddRequestFailure TODO
func memberRequestFailure(err error, groupName string, logger *zap.SugaredLogger, memberName, memberType string) (resp *api.GroupMembersAddDefault) {

	// Log the error.
	message := fmt.Sprintf("Failed to add member \"%s\" to the quota-group.", memberName)
	logger.Infow(message,
		"quota-group", groupName,
		"memberName", memberName,
		"memberType", memberType,
		"error", err.Error(),
	)

	// Create the response back to the client.
	code := 500
	resp = &api.GroupMembersAddDefault{Payload: &models.Error{
		Code:    int64(code),
		Message: message,
	}}
	resp.SetStatusCode(code)

	return resp
}
