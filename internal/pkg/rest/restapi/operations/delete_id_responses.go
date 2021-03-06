// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	models "payments-simple/internal/pkg/rest/models"

	"github.com/go-openapi/runtime"
)

// DeleteIDOKCode is the HTTP code returned for type DeleteIDOK
const DeleteIDOKCode int = 200

/*DeleteIDOK success

swagger:response deleteIdOK
*/
type DeleteIDOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Payment `json:"body,omitempty"`
}

// NewDeleteIDOK creates DeleteIDOK with default headers values
func NewDeleteIDOK() *DeleteIDOK {

	return &DeleteIDOK{}
}

// WithPayload adds the payload to the delete Id o k response
func (o *DeleteIDOK) WithPayload(payload []*models.Payment) *DeleteIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Id o k response
func (o *DeleteIDOK) SetPayload(payload []*models.Payment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Payment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*DeleteIDDefault generic error

swagger:response deleteIdDefault
*/
type DeleteIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewDeleteIDDefault creates DeleteIDDefault with default headers values
func NewDeleteIDDefault(code int) *DeleteIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete ID default response
func (o *DeleteIDDefault) WithStatusCode(code int) *DeleteIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete ID default response
func (o *DeleteIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete ID default response
func (o *DeleteIDDefault) WithPayload(payload *models.Message) *DeleteIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete ID default response
func (o *DeleteIDDefault) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
