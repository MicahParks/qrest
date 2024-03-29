// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/qrest/models"
)

// GroupMembersAddOKCode is the HTTP code returned for type GroupMembersAddOK
const GroupMembersAddOKCode int = 200

/*GroupMembersAddOK The members of the group were successfully added.

swagger:response groupMembersAddOK
*/
type GroupMembersAddOK struct {
}

// NewGroupMembersAddOK creates GroupMembersAddOK with default headers values
func NewGroupMembersAddOK() *GroupMembersAddOK {

	return &GroupMembersAddOK{}
}

// WriteResponse to the client
func (o *GroupMembersAddOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*GroupMembersAddDefault Unexpected error.

swagger:response groupMembersAddDefault
*/
type GroupMembersAddDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGroupMembersAddDefault creates GroupMembersAddDefault with default headers values
func NewGroupMembersAddDefault(code int) *GroupMembersAddDefault {
	if code <= 0 {
		code = 500
	}

	return &GroupMembersAddDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the group members add default response
func (o *GroupMembersAddDefault) WithStatusCode(code int) *GroupMembersAddDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the group members add default response
func (o *GroupMembersAddDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the group members add default response
func (o *GroupMembersAddDefault) WithPayload(payload *models.Error) *GroupMembersAddDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the group members add default response
func (o *GroupMembersAddDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GroupMembersAddDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
