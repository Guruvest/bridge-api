// Code generated by go-swagger; DO NOT EDIT.

package account_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPortfoliosByAccountParams creates a new GetPortfoliosByAccountParams object
// with the default values initialized.
func NewGetPortfoliosByAccountParams() GetPortfoliosByAccountParams {
	var ()
	return GetPortfoliosByAccountParams{}
}

// GetPortfoliosByAccountParams contains all the bound params for the get portfolios by account operation
// typically these are obtained from a http.Request
//
// swagger:parameters getPortfoliosByAccount
type GetPortfoliosByAccountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The linked account identifier
	  In: query
	*/
	AccountID *int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetPortfoliosByAccountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAccountID, qhkAccountID, _ := qs.GetOK("account_id")
	if err := o.bindAccountID(qAccountID, qhkAccountID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPortfoliosByAccountParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("account_id", "query", "int32", raw)
	}
	o.AccountID = &value

	return nil
}