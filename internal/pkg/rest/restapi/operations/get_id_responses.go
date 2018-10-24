// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	models "payments-simple/internal/pkg/rest/models"

	"github.com/go-openapi/runtime"
)

// GetIDOKCode is the HTTP code returned for type GetIDOK
const GetIDOKCode int = 200

/*GetIDOK success

swagger:response getIdOK
*/
type GetIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Payment `json:"body,omitempty"`
}

// NewGetIDOK creates GetIDOK with default headers values
func NewGetIDOK() *GetIDOK {

	return &GetIDOK{}
}

// WithPayload adds the payload to the get Id o k response
func (o *GetIDOK) WithPayload(payload *models.Payment) *GetIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Id o k response
func (o *GetIDOK) SetPayload(payload *models.Payment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetIDDefault generic error

swagger:response getIdDefault
*/
type GetIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewGetIDDefault creates GetIDDefault with default headers values
func NewGetIDDefault(code int) *GetIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get ID default response
func (o *GetIDDefault) WithStatusCode(code int) *GetIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get ID default response
func (o *GetIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get ID default response
func (o *GetIDDefault) WithPayload(payload *models.Message) *GetIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ID default response
func (o *GetIDDefault) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}