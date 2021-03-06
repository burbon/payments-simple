// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	models "payments-simple/internal/pkg/rest/models"

	"github.com/go-openapi/runtime"
)

// GetPingOKCode is the HTTP code returned for type GetPingOK
const GetPingOKCode int = 200

/*GetPingOK success

swagger:response getPingOK
*/
type GetPingOK struct {

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewGetPingOK creates GetPingOK with default headers values
func NewGetPingOK() *GetPingOK {

	return &GetPingOK{}
}

// WithPayload adds the payload to the get ping o k response
func (o *GetPingOK) WithPayload(payload *models.Message) *GetPingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ping o k response
func (o *GetPingOK) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetPingDefault generic error

swagger:response getPingDefault
*/
type GetPingDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewGetPingDefault creates GetPingDefault with default headers values
func NewGetPingDefault(code int) *GetPingDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPingDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get ping default response
func (o *GetPingDefault) WithStatusCode(code int) *GetPingDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get ping default response
func (o *GetPingDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get ping default response
func (o *GetPingDefault) WithPayload(payload *models.Message) *GetPingDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ping default response
func (o *GetPingDefault) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPingDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
