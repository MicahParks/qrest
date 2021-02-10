// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GroupMembersReadHandlerFunc turns a function with the right signature into a group members read handler
type GroupMembersReadHandlerFunc func(GroupMembersReadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GroupMembersReadHandlerFunc) Handle(params GroupMembersReadParams) middleware.Responder {
	return fn(params)
}

// GroupMembersReadHandler interface for that can handle valid group members read params
type GroupMembersReadHandler interface {
	Handle(GroupMembersReadParams) middleware.Responder
}

// NewGroupMembersRead creates a new http.Handler for the group members read operation
func NewGroupMembersRead(ctx *middleware.Context, handler GroupMembersReadHandler) *GroupMembersRead {
	return &GroupMembersRead{Context: ctx, Handler: handler}
}

/* GroupMembersRead swagger:route GET /group/members api groupMembersRead

Get the members of quota-groups.

*/
type GroupMembersRead struct {
	Context *middleware.Context
	Handler GroupMembersReadHandler
}

func (o *GroupMembersRead) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGroupMembersReadParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}