// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GroupMembersDeleteHandlerFunc turns a function with the right signature into a group members delete handler
type GroupMembersDeleteHandlerFunc func(GroupMembersDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GroupMembersDeleteHandlerFunc) Handle(params GroupMembersDeleteParams) middleware.Responder {
	return fn(params)
}

// GroupMembersDeleteHandler interface for that can handle valid group members delete params
type GroupMembersDeleteHandler interface {
	Handle(GroupMembersDeleteParams) middleware.Responder
}

// NewGroupMembersDelete creates a new http.Handler for the group members delete operation
func NewGroupMembersDelete(ctx *middleware.Context, handler GroupMembersDeleteHandler) *GroupMembersDelete {
	return &GroupMembersDelete{Context: ctx, Handler: handler}
}

/* GroupMembersDelete swagger:route DELETE /group/members api groupMembersDelete

Delete members from the quota-group.

Given a map of quota-groups to members, disassociate the given members from their associated quota-groups on the backend.

*/
type GroupMembersDelete struct {
	Context *middleware.Context
	Handler GroupMembersDeleteHandler
}

func (o *GroupMembersDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGroupMembersDeleteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
