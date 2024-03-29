// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GroupDeleteHandlerFunc turns a function with the right signature into a group delete handler
type GroupDeleteHandlerFunc func(GroupDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GroupDeleteHandlerFunc) Handle(params GroupDeleteParams) middleware.Responder {
	return fn(params)
}

// GroupDeleteHandler interface for that can handle valid group delete params
type GroupDeleteHandler interface {
	Handle(GroupDeleteParams) middleware.Responder
}

// NewGroupDelete creates a new http.Handler for the group delete operation
func NewGroupDelete(ctx *middleware.Context, handler GroupDeleteHandler) *GroupDelete {
	return &GroupDelete{Context: ctx, Handler: handler}
}

/* GroupDelete swagger:route DELETE /group api groupDelete

Delete quota-groups.

Delete all the given quota-groups from the backend.

*/
type GroupDelete struct {
	Context *middleware.Context
	Handler GroupDeleteHandler
}

func (o *GroupDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGroupDeleteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
