// Code generated by go-swagger; DO NOT EDIT.

package market_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateOrderHandlerFunc turns a function with the right signature into a update order handler
type UpdateOrderHandlerFunc func(UpdateOrderParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateOrderHandlerFunc) Handle(params UpdateOrderParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateOrderHandler interface for that can handle valid update order params
type UpdateOrderHandler interface {
	Handle(UpdateOrderParams, interface{}) middleware.Responder
}

// NewUpdateOrder creates a new http.Handler for the update order operation
func NewUpdateOrder(ctx *middleware.Context, handler UpdateOrderHandler) *UpdateOrder {
	return &UpdateOrder{Context: ctx, Handler: handler}
}

/*UpdateOrder swagger:route PUT /market/orders/{order_id} Market Access updateOrder

Update an existing Order

*/
type UpdateOrder struct {
	Context *middleware.Context
	Handler UpdateOrderHandler
}

func (o *UpdateOrder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateOrderParams()

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
