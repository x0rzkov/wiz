// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetConfigOKCode is the HTTP code returned for type GetConfigOK
const GetConfigOKCode int = 200

/*GetConfigOK OK

swagger:response getConfigOK
*/
type GetConfigOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetConfigOK creates GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {

	return &GetConfigOK{}
}

// WithPayload adds the payload to the get config o k response
func (o *GetConfigOK) WithPayload(payload interface{}) *GetConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config o k response
func (o *GetConfigOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetConfigInternalServerErrorCode is the HTTP code returned for type GetConfigInternalServerError
const GetConfigInternalServerErrorCode int = 500

/*GetConfigInternalServerError Internal Server Error

swagger:response getConfigInternalServerError
*/
type GetConfigInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetConfigInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetConfigInternalServerError creates GetConfigInternalServerError with default headers values
func NewGetConfigInternalServerError() *GetConfigInternalServerError {

	return &GetConfigInternalServerError{}
}

// WithPayload adds the payload to the get config internal server error response
func (o *GetConfigInternalServerError) WithPayload(payload *GetConfigInternalServerErrorBody) *GetConfigInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config internal server error response
func (o *GetConfigInternalServerError) SetPayload(payload *GetConfigInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
