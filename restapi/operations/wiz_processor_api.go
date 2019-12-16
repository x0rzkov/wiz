// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/alexkreidler/wiz/restapi/operations/processors"
)

// NewWizProcessorAPI creates a new WizProcessor instance
func NewWizProcessorAPI(spec *loads.Document) *WizProcessorAPI {
	return &WizProcessorAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		ProcessorsAddDataHandler: processors.AddDataHandlerFunc(func(params processors.AddDataParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsAddData has not yet been implemented")
		}),
		ProcessorsGetAllProcessorsHandler: processors.GetAllProcessorsHandlerFunc(func(params processors.GetAllProcessorsParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsGetAllProcessors has not yet been implemented")
		}),
		ProcessorsGetAllRunsHandler: processors.GetAllRunsHandlerFunc(func(params processors.GetAllRunsParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsGetAllRuns has not yet been implemented")
		}),
		ProcessorsGetConfigHandler: processors.GetConfigHandlerFunc(func(params processors.GetConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsGetConfig has not yet been implemented")
		}),
		ProcessorsGetDataHandler: processors.GetDataHandlerFunc(func(params processors.GetDataParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsGetData has not yet been implemented")
		}),
		ProcessorsGetInputChunkHandler: processors.GetInputChunkHandlerFunc(func(params processors.GetInputChunkParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsGetInputChunk has not yet been implemented")
		}),
		ProcessorsGetOutputChunkHandler: processors.GetOutputChunkHandlerFunc(func(params processors.GetOutputChunkParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsGetOutputChunk has not yet been implemented")
		}),
		ProcessorsGetProcessorHandler: processors.GetProcessorHandlerFunc(func(params processors.GetProcessorParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsGetProcessor has not yet been implemented")
		}),
		ProcessorsGetRunHandler: processors.GetRunHandlerFunc(func(params processors.GetRunParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsGetRun has not yet been implemented")
		}),
		ProcessorsUpdateConfigHandler: processors.UpdateConfigHandlerFunc(func(params processors.UpdateConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation ProcessorsUpdateConfig has not yet been implemented")
		}),
	}
}

/*WizProcessorAPI This document defines the Wiz Processor API as implemented using JSON over HTTP. */
type WizProcessorAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ProcessorsAddDataHandler sets the operation handler for the add data operation
	ProcessorsAddDataHandler processors.AddDataHandler
	// ProcessorsGetAllProcessorsHandler sets the operation handler for the get all processors operation
	ProcessorsGetAllProcessorsHandler processors.GetAllProcessorsHandler
	// ProcessorsGetAllRunsHandler sets the operation handler for the get all runs operation
	ProcessorsGetAllRunsHandler processors.GetAllRunsHandler
	// ProcessorsGetConfigHandler sets the operation handler for the get config operation
	ProcessorsGetConfigHandler processors.GetConfigHandler
	// ProcessorsGetDataHandler sets the operation handler for the get data operation
	ProcessorsGetDataHandler processors.GetDataHandler
	// ProcessorsGetInputChunkHandler sets the operation handler for the get input chunk operation
	ProcessorsGetInputChunkHandler processors.GetInputChunkHandler
	// ProcessorsGetOutputChunkHandler sets the operation handler for the get output chunk operation
	ProcessorsGetOutputChunkHandler processors.GetOutputChunkHandler
	// ProcessorsGetProcessorHandler sets the operation handler for the get processor operation
	ProcessorsGetProcessorHandler processors.GetProcessorHandler
	// ProcessorsGetRunHandler sets the operation handler for the get run operation
	ProcessorsGetRunHandler processors.GetRunHandler
	// ProcessorsUpdateConfigHandler sets the operation handler for the update config operation
	ProcessorsUpdateConfigHandler processors.UpdateConfigHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *WizProcessorAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WizProcessorAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WizProcessorAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WizProcessorAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WizProcessorAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WizProcessorAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WizProcessorAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WizProcessorAPI
func (o *WizProcessorAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ProcessorsAddDataHandler == nil {
		unregistered = append(unregistered, "processors.AddDataHandler")
	}

	if o.ProcessorsGetAllProcessorsHandler == nil {
		unregistered = append(unregistered, "processors.GetAllProcessorsHandler")
	}

	if o.ProcessorsGetAllRunsHandler == nil {
		unregistered = append(unregistered, "processors.GetAllRunsHandler")
	}

	if o.ProcessorsGetConfigHandler == nil {
		unregistered = append(unregistered, "processors.GetConfigHandler")
	}

	if o.ProcessorsGetDataHandler == nil {
		unregistered = append(unregistered, "processors.GetDataHandler")
	}

	if o.ProcessorsGetInputChunkHandler == nil {
		unregistered = append(unregistered, "processors.GetInputChunkHandler")
	}

	if o.ProcessorsGetOutputChunkHandler == nil {
		unregistered = append(unregistered, "processors.GetOutputChunkHandler")
	}

	if o.ProcessorsGetProcessorHandler == nil {
		unregistered = append(unregistered, "processors.GetProcessorHandler")
	}

	if o.ProcessorsGetRunHandler == nil {
		unregistered = append(unregistered, "processors.GetRunHandler")
	}

	if o.ProcessorsUpdateConfigHandler == nil {
		unregistered = append(unregistered, "processors.UpdateConfigHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WizProcessorAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WizProcessorAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *WizProcessorAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *WizProcessorAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *WizProcessorAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *WizProcessorAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the wiz processor API
func (o *WizProcessorAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WizProcessorAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/processor/{id}/run/{runID}/data/input/{chunkID}"] = processors.NewAddData(o.context, o.ProcessorsAddDataHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/processors"] = processors.NewGetAllProcessors(o.context, o.ProcessorsGetAllProcessorsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/processor/{id}/runs"] = processors.NewGetAllRuns(o.context, o.ProcessorsGetAllRunsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/processor/{id}/run/{runID}/configuration"] = processors.NewGetConfig(o.context, o.ProcessorsGetConfigHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/processor/{id}/run/{runID}/data"] = processors.NewGetData(o.context, o.ProcessorsGetDataHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/processor/{id}/run/{runID}/data/input/{chunkID}"] = processors.NewGetInputChunk(o.context, o.ProcessorsGetInputChunkHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/processor/{id}/run/{runID}/data/output/{chunkID}"] = processors.NewGetOutputChunk(o.context, o.ProcessorsGetOutputChunkHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/processor/{id}"] = processors.NewGetProcessor(o.context, o.ProcessorsGetProcessorHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/processor/{id}/run/{runID}"] = processors.NewGetRun(o.context, o.ProcessorsGetRunHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/processor/{id}/run/{runID}/configuration"] = processors.NewUpdateConfig(o.context, o.ProcessorsUpdateConfigHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WizProcessorAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *WizProcessorAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *WizProcessorAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *WizProcessorAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
