// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupMembersAddHandlerFunc turns a function with the right signature into a group members add handler
type GroupMembersAddHandlerFunc func(GroupMembersAddParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GroupMembersAddHandlerFunc) Handle(params GroupMembersAddParams) middleware.Responder {
	return fn(params)
}

// GroupMembersAddHandler interface for that can handle valid group members add params
type GroupMembersAddHandler interface {
	Handle(GroupMembersAddParams) middleware.Responder
}

// NewGroupMembersAdd creates a new http.Handler for the group members add operation
func NewGroupMembersAdd(ctx *middleware.Context, handler GroupMembersAddHandler) *GroupMembersAdd {
	return &GroupMembersAdd{Context: ctx, Handler: handler}
}

/* GroupMembersAdd swagger:route POST /group/{group}/members groupMembersAdd

Add members to the quota-group.

*/
type GroupMembersAdd struct {
	Context *middleware.Context
	Handler GroupMembersAddHandler
}

func (o *GroupMembersAdd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGroupMembersAddParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GroupMembersAddBody group members add body
//
// swagger:model GroupMembersAddBody
type GroupMembersAddBody struct {

	// snaps
	Snaps []string `json:"snaps"`

	// sub groups
	SubGroups []string `json:"subGroups"`
}

// Validate validates this group members add body
func (o *GroupMembersAddBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this group members add body based on context it is used
func (o *GroupMembersAddBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GroupMembersAddBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GroupMembersAddBody) UnmarshalBinary(b []byte) error {
	var res GroupMembersAddBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}