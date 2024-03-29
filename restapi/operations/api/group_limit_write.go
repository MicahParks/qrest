// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GroupLimitWriteHandlerFunc turns a function with the right signature into a group limit write handler
type GroupLimitWriteHandlerFunc func(GroupLimitWriteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GroupLimitWriteHandlerFunc) Handle(params GroupLimitWriteParams) middleware.Responder {
	return fn(params)
}

// GroupLimitWriteHandler interface for that can handle valid group limit write params
type GroupLimitWriteHandler interface {
	Handle(GroupLimitWriteParams) middleware.Responder
}

// NewGroupLimitWrite creates a new http.Handler for the group limit write operation
func NewGroupLimitWrite(ctx *middleware.Context, handler GroupLimitWriteHandler) *GroupLimitWrite {
	return &GroupLimitWrite{Context: ctx, Handler: handler}
}

/* GroupLimitWrite swagger:route POST /group/limits api groupLimitWrite

Set the resource limits for the given quota-groups.

Given a map of quota-group names to resource limits, set the resource limits for the quota-groups on the backend.

*/
type GroupLimitWrite struct {
	Context *middleware.Context
	Handler GroupLimitWriteHandler
}

func (o *GroupLimitWrite) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGroupLimitWriteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
