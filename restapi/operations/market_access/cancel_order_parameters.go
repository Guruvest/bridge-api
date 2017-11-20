// Code generated by go-swagger; DO NOT EDIT.

package market_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCancelOrderParams creates a new CancelOrderParams object
// with the default values initialized.
func NewCancelOrderParams() CancelOrderParams {
	var ()
	return CancelOrderParams{}
}

// CancelOrderParams contains all the bound params for the cancel order operation
// typically these are obtained from a http.Request
//
// swagger:parameters cancelOrder
type CancelOrderParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The order identifier
	  Required: true
	  In: path
	*/
	OrderID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CancelOrderParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rOrderID, rhkOrderID, _ := route.Params.GetOK("order_id")
	if err := o.bindOrderID(rOrderID, rhkOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelOrderParams) bindOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.OrderID = raw

	return nil
}
