// Code generated by go-swagger; DO NOT EDIT.

package account_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/alien45/bridge-api/models"
)

// GetTradesByAccountOKCode is the HTTP code returned for type GetTradesByAccountOK
const GetTradesByAccountOKCode int = 200

/*GetTradesByAccountOK successful operation

swagger:response getTradesByAccountOK
*/
type GetTradesByAccountOK struct {

	/*
	  In: Body
	*/
	Payload models.GetTradesByAccountOKBody `json:"body,omitempty"`
}

// NewGetTradesByAccountOK creates GetTradesByAccountOK with default headers values
func NewGetTradesByAccountOK() *GetTradesByAccountOK {
	return &GetTradesByAccountOK{}
}

// WithPayload adds the payload to the get trades by account o k response
func (o *GetTradesByAccountOK) WithPayload(payload models.GetTradesByAccountOKBody) *GetTradesByAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get trades by account o k response
func (o *GetTradesByAccountOK) SetPayload(payload models.GetTradesByAccountOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTradesByAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetTradesByAccountOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetTradesByAccountDefault unexpected error

swagger:response getTradesByAccountDefault
*/
type GetTradesByAccountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetTradesByAccountDefault creates GetTradesByAccountDefault with default headers values
func NewGetTradesByAccountDefault(code int) *GetTradesByAccountDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTradesByAccountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get trades by account default response
func (o *GetTradesByAccountDefault) WithStatusCode(code int) *GetTradesByAccountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get trades by account default response
func (o *GetTradesByAccountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get trades by account default response
func (o *GetTradesByAccountDefault) WithPayload(payload *models.ErrorResponse) *GetTradesByAccountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get trades by account default response
func (o *GetTradesByAccountDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTradesByAccountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
