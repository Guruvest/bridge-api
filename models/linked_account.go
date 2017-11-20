// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LinkedAccount Information about a linked account
// swagger:model LinkedAccount
type LinkedAccount struct {

	// Identify if the account is active or not
	Active bool `json:"active,omitempty"`

	// Guruvest internal exchange/broker identifier
	BrokerageID int32 `json:"brokerage_id,omitempty"`

	// The name of the Venue (broker or exchange)
	BrokerageName string `json:"brokerage_name,omitempty"`

	// The unique Guruvest internal identifier for the linked account
	ID int32 `json:"id,omitempty"`

	// Identify if the account should be used as the default one
	IsDefault bool `json:"is_default,omitempty"`

	// The name of the account
	Name string `json:"name,omitempty"`

	// Identify if the account is able to perform trades or not
	Tradable bool `json:"tradable,omitempty"`
}

// Validate validates this linked account
func (m *LinkedAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *LinkedAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkedAccount) UnmarshalBinary(b []byte) error {
	var res LinkedAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
