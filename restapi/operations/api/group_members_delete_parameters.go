// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/MicahParks/qrest/models"
)

// NewGroupMembersDeleteParams creates a new GroupMembersDeleteParams object
//
// There are no default values defined in the spec.
func NewGroupMembersDeleteParams() GroupMembersDeleteParams {

	return GroupMembersDeleteParams{}
}

// GroupMembersDeleteParams contains all the bound params for the group members delete operation
// typically these are obtained from a http.Request
//
// swagger:parameters groupMembersDelete
type GroupMembersDeleteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The mapping of quota-group names to the snaps and member quota-groups to remove.
	  In: body
	*/
	GroupMembersMap map[string]models.GroupMembers
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGroupMembersDeleteParams() beforehand.
func (o *GroupMembersDeleteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body map[string]models.GroupMembers
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("groupMembersMap", "body", "", err))
		} else {
			// validate map of body objects
			for k := range body {
				if err := validate.Required(fmt.Sprintf("%s.%v", "groupMembersMap", k), "body", body[k]); err != nil {
					return err
				}
				if val, ok := body[k]; ok {
					if err := val.Validate(route.Formats); err != nil {
						res = append(res, err)
						break
					}
				}
			}

			if len(res) == 0 {
				o.GroupMembersMap = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
