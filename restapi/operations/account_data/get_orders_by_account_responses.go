// Code generated by go-swagger; DO NOT EDIT.

package account_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/alien45/bridge-api/models"
)

// GetOrdersByAccountOKCode is the HTTP code returned for type GetOrdersByAccountOK
const GetOrdersByAccountOKCode int = 200

/*GetOrdersByAccountOK successful operation

swagger:response getOrdersByAccountOK
*/
type GetOrdersByAccountOK struct {

	/*
	  In: Body
	*/
	Payload models.GetOrdersByAccountOKBody `json:"body,omitempty"`
}

// NewGetOrdersByAccountOK creates GetOrdersByAccountOK with default headers values
func NewGetOrdersByAccountOK() *GetOrdersByAccountOK {
	return &GetOrdersByAccountOK{}
}

// WithPayload adds the payload to the get orders by account o k response
func (o *GetOrdersByAccountOK) WithPayload(payload models.GetOrdersByAccountOKBody) *GetOrdersByAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get orders by account o k response
func (o *GetOrdersByAccountOK) SetPayload(payload models.GetOrdersByAccountOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrdersByAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetOrdersByAccountOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetOrdersByAccountDefault unexpected error

swagger:response getOrdersByAccountDefault
*/
type GetOrdersByAccountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetOrdersByAccountDefault creates GetOrdersByAccountDefault with default headers values
func NewGetOrdersByAccountDefault(code int) *GetOrdersByAccountDefault {
	if code <= 0 {
		code = 500
	}

	return &GetOrdersByAccountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get orders by account default response
func (o *GetOrdersByAccountDefault) WithStatusCode(code int) *GetOrdersByAccountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get orders by account default response
func (o *GetOrdersByAccountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get orders by account default response
func (o *GetOrdersByAccountDefault) WithPayload(payload *models.ErrorResponse) *GetOrdersByAccountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get orders by account default response
func (o *GetOrdersByAccountDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrdersByAccountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
