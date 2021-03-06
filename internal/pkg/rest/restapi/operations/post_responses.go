// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	models "payments-simple/internal/pkg/rest/models"

	"github.com/go-openapi/runtime"
)

// PostCreatedCode is the HTTP code returned for type PostCreated
const PostCreatedCode int = 201

/*PostCreated created

swagger:response postCreated
*/
type PostCreated struct {
}

// NewPostCreated creates PostCreated with default headers values
func NewPostCreated() *PostCreated {

	return &PostCreated{}
}

// WriteResponse to the client
func (o *PostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*PostDefault generic error

swagger:response postDefault
*/
type PostDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewPostDefault creates PostDefault with default headers values
func NewPostDefault(code int) *PostDefault {
	if code <= 0 {
		code = 500
	}

	return &PostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post default response
func (o *PostDefault) WithStatusCode(code int) *PostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post default response
func (o *PostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post default response
func (o *PostDefault) WithPayload(payload *models.Message) *PostDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post default response
func (o *PostDefault) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
