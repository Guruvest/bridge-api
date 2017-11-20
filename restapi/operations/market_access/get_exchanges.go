// Code generated by go-swagger; DO NOT EDIT.

package market_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetExchangesHandlerFunc turns a function with the right signature into a get exchanges handler
type GetExchangesHandlerFunc func(GetExchangesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetExchangesHandlerFunc) Handle(params GetExchangesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetExchangesHandler interface for that can handle valid get exchanges params
type GetExchangesHandler interface {
	Handle(GetExchangesParams, interface{}) middleware.Responder
}

// NewGetExchanges creates a new http.Handler for the get exchanges operation
func NewGetExchanges(ctx *middleware.Context, handler GetExchangesHandler) *GetExchanges {
	return &GetExchanges{Context: ctx, Handler: handler}
}

/*GetExchanges swagger:route GET /market/exchanges Market Access getExchanges

Return the list of exchanges supported

*/
type GetExchanges struct {
	Context *middleware.Context
	Handler GetExchangesHandler
}

func (o *GetExchanges) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetExchangesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
