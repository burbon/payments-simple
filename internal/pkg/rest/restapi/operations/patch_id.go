// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchIDHandlerFunc turns a function with the right signature into a patch ID handler
type PatchIDHandlerFunc func(PatchIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchIDHandlerFunc) Handle(params PatchIDParams) middleware.Responder {
	return fn(params)
}

// PatchIDHandler interface for that can handle valid patch ID params
type PatchIDHandler interface {
	Handle(PatchIDParams) middleware.Responder
}

// NewPatchID creates a new http.Handler for the patch ID operation
func NewPatchID(ctx *middleware.Context, handler PatchIDHandler) *PatchID {
	return &PatchID{Context: ctx, Handler: handler}
}

/*PatchID swagger:route PATCH /{id} patchId

patch payment

*/
type PatchID struct {
	Context *middleware.Context
	Handler PatchIDHandler
}

func (o *PatchID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
