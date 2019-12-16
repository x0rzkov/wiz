// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// GetDataHandlerFunc turns a function with the right signature into a get data handler
type GetDataHandlerFunc func(GetDataParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDataHandlerFunc) Handle(params GetDataParams) middleware.Responder {
	return fn(params)
}

// GetDataHandler interface for that can handle valid get data params
type GetDataHandler interface {
	Handle(GetDataParams) middleware.Responder
}

// NewGetData creates a new http.Handler for the get data operation
func NewGetData(ctx *middleware.Context, handler GetDataHandler) *GetData {
	return &GetData{Context: ctx, Handler: handler}
}

/*GetData swagger:route GET /processor/{id}/run/{runID}/data Processors getData

GetData

Returns all the data chunks which the processor has recieved, and all that it has written

TODO: figure out if each chunk should report state: yes

*/
type GetData struct {
	Context *middleware.Context
	Handler GetDataHandler
}

func (o *GetData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDataParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetDataOKBody get data o k body
// swagger:model GetDataOKBody
type GetDataOKBody struct {

	// in
	In []interface{} `json:"in"`

	// out
	Out []interface{} `json:"out"`
}

// Validate validates this get data o k body
func (o *GetDataOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDataOKBody) UnmarshalBinary(b []byte) error {
	var res GetDataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
