// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mvo5/qrest-skeleton/models"
)

// GroupMembersDeleteOKCode is the HTTP code returned for type GroupMembersDeleteOK
const GroupMembersDeleteOKCode int = 200

/*GroupMembersDeleteOK The members of the group were successfully deleted.

swagger:response groupMembersDeleteOK
*/
type GroupMembersDeleteOK struct {
}

// NewGroupMembersDeleteOK creates GroupMembersDeleteOK with default headers values
func NewGroupMembersDeleteOK() *GroupMembersDeleteOK {

	return &GroupMembersDeleteOK{}
}

// WriteResponse to the client
func (o *GroupMembersDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*GroupMembersDeleteDefault Unexpected error.

swagger:response groupMembersDeleteDefault
*/
type GroupMembersDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGroupMembersDeleteDefault creates GroupMembersDeleteDefault with default headers values
func NewGroupMembersDeleteDefault(code int) *GroupMembersDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &GroupMembersDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the group members delete default response
func (o *GroupMembersDeleteDefault) WithStatusCode(code int) *GroupMembersDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the group members delete default response
func (o *GroupMembersDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the group members delete default response
func (o *GroupMembersDeleteDefault) WithPayload(payload *models.Error) *GroupMembersDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the group members delete default response
func (o *GroupMembersDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GroupMembersDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}