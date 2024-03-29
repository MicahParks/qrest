// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GroupLimitReadHandlerFunc turns a function with the right signature into a group limit read handler
type GroupLimitReadHandlerFunc func(GroupLimitReadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GroupLimitReadHandlerFunc) Handle(params GroupLimitReadParams) middleware.Responder {
	return fn(params)
}

// GroupLimitReadHandler interface for that can handle valid group limit read params
type GroupLimitReadHandler interface {
	Handle(GroupLimitReadParams) middleware.Responder
}

// NewGroupLimitRead creates a new http.Handler for the group limit read operation
func NewGroupLimitRead(ctx *middleware.Context, handler GroupLimitReadHandler) *GroupLimitRead {
	return &GroupLimitRead{Context: ctx, Handler: handler}
}

/* GroupLimitRead swagger:route GET /group/limits api groupLimitRead

Get the resource limits for the given quota-groups.

Given an array of quota-groups, return all of their limits in a map.

*/
type GroupLimitRead struct {
	Context *middleware.Context
	Handler GroupLimitReadHandler
}

func (o *GroupLimitRead) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGroupLimitReadParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
