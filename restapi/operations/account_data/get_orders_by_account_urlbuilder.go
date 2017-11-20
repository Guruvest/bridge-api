// Code generated by go-swagger; DO NOT EDIT.

package account_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetOrdersByAccountURL generates an URL for the get orders by account operation
type GetOrdersByAccountURL struct {
	AccountID *int32
	Limit     *int32
	Status    *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetOrdersByAccountURL) WithBasePath(bp string) *GetOrdersByAccountURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetOrdersByAccountURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetOrdersByAccountURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/accounts/orders"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var accountID string
	if o.AccountID != nil {
		accountID = swag.FormatInt32(*o.AccountID)
	}
	if accountID != "" {
		qs.Set("account_id", accountID)
	}

	var limit string
	if o.Limit != nil {
		limit = swag.FormatInt32(*o.Limit)
	}
	if limit != "" {
		qs.Set("limit", limit)
	}

	var status string
	if o.Status != nil {
		status = *o.Status
	}
	if status != "" {
		qs.Set("status", status)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetOrdersByAccountURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetOrdersByAccountURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetOrdersByAccountURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetOrdersByAccountURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetOrdersByAccountURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetOrdersByAccountURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
