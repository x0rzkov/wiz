// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// GetInputChunkHandlerFunc turns a function with the right signature into a get input chunk handler
type GetInputChunkHandlerFunc func(GetInputChunkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInputChunkHandlerFunc) Handle(params GetInputChunkParams) middleware.Responder {
	return fn(params)
}

// GetInputChunkHandler interface for that can handle valid get input chunk params
type GetInputChunkHandler interface {
	Handle(GetInputChunkParams) middleware.Responder
}

// NewGetInputChunk creates a new http.Handler for the get input chunk operation
func NewGetInputChunk(ctx *middleware.Context, handler GetInputChunkHandler) *GetInputChunk {
	return &GetInputChunk{Context: ctx, Handler: handler}
}

/*GetInputChunk swagger:route GET /processor/{id}/run/{runID}/data/input/{chunkID} Processors getInputChunk

GetInputChunk

This is not likely to be used, as we might as well get all chunks, but it needs to be here.

*/
type GetInputChunk struct {
	Context *middleware.Context
	Handler GetInputChunkHandler
}

func (o *GetInputChunk) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetInputChunkParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetInputChunkOKBody get input chunk o k body
// swagger:model GetInputChunkOKBody
type GetInputChunkOKBody struct {

	// the ID of the chunk
	// Required: true
	ChunkID *string `json:"chunk_id"`

	// error
	Error *GetInputChunkOKBodyError `json:"error,omitempty"`

	// the output chunk ID association
	OutputChunkID string `json:"output_chunk_id,omitempty"`

	// the state of the data chunk.
	// Required: true
	// Enum: [Syncing Determining Validating Running Failed Succeeded]
	State *string `json:"state"`
}

// Validate validates this get input chunk o k body
func (o *GetInputChunkOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChunkID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInputChunkOKBody) validateChunkID(formats strfmt.Registry) error {

	if err := validate.Required("getInputChunkOK"+"."+"chunk_id", "body", o.ChunkID); err != nil {
		return err
	}

	return nil
}

func (o *GetInputChunkOKBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInputChunkOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

var getInputChunkOKBodyTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Syncing","Determining","Validating","Running","Failed","Succeeded"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getInputChunkOKBodyTypeStatePropEnum = append(getInputChunkOKBodyTypeStatePropEnum, v)
	}
}

const (

	// GetInputChunkOKBodyStateSyncing captures enum value "Syncing"
	GetInputChunkOKBodyStateSyncing string = "Syncing"

	// GetInputChunkOKBodyStateDetermining captures enum value "Determining"
	GetInputChunkOKBodyStateDetermining string = "Determining"

	// GetInputChunkOKBodyStateValidating captures enum value "Validating"
	GetInputChunkOKBodyStateValidating string = "Validating"

	// GetInputChunkOKBodyStateRunning captures enum value "Running"
	GetInputChunkOKBodyStateRunning string = "Running"

	// GetInputChunkOKBodyStateFailed captures enum value "Failed"
	GetInputChunkOKBodyStateFailed string = "Failed"

	// GetInputChunkOKBodyStateSucceeded captures enum value "Succeeded"
	GetInputChunkOKBodyStateSucceeded string = "Succeeded"
)

// prop value enum
func (o *GetInputChunkOKBody) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getInputChunkOKBodyTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetInputChunkOKBody) validateState(formats strfmt.Registry) error {

	if err := validate.Required("getInputChunkOK"+"."+"state", "body", o.State); err != nil {
		return err
	}

	// value enum
	if err := o.validateStateEnum("getInputChunkOK"+"."+"state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInputChunkOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInputChunkOKBody) UnmarshalBinary(b []byte) error {
	var res GetInputChunkOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetInputChunkOKBodyError the error which caused the chunk to fail if it is in the Failed state
// swagger:model GetInputChunkOKBodyError
type GetInputChunkOKBodyError struct {

	// The full trace of the error. This may only be available in a debug mode
	Trace string `json:"trace,omitempty"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this get input chunk o k body error
func (o *GetInputChunkOKBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInputChunkOKBodyError) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("getInputChunkOK"+"."+"error"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInputChunkOKBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInputChunkOKBodyError) UnmarshalBinary(b []byte) error {
	var res GetInputChunkOKBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
