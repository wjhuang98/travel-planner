// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version v0.3.0 DO NOT EDIT.
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/discord-gophers/goapi-gen/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Place defines model for Place.
type Place struct {
	Address string   `json:"address"`
	Name    string   `json:"name"`
	Photos  []string `json:"photos"`
	Rating  float32  `json:"rating"`
}

// PlacesList defines model for PlacesList.
type PlacesList []Place

// SearchParams defines parameters for Search.
type SearchParams struct {
	// Hotels, Restaurants, or Attractions
	Filter string `json:"filter"`

	// City Name
	Location string  `json:"location"`
	Distance float64 `json:"distance"`
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
// It may also be instantiated directly, for the purpose of responding with a single status code.
type Response struct {
	body        interface{}
	Code        int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.Code)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(code int) *Response {
	resp.Code = code
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// SearchJSON200Response is a constructor method for a Search response.
// A *Response is returned with the configured status code and content type from the spec.
func SearchJSON200Response(body PlacesList) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// SearchJSON400Response is a constructor method for a Search response.
// A *Response is returned with the configured status code and content type from the spec.
func SearchJSON400Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        400,
		contentType: "application/json",
	}
}

// SearchJSON500Response is a constructor method for a Search response.
// A *Response is returned with the configured status code and content type from the spec.
func SearchJSON500Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        500,
		contentType: "application/json",
	}
}

// SearchJSON502Response is a constructor method for a Search response.
// A *Response is returned with the configured status code and content type from the spec.
func SearchJSON502Response(body Error) *Response {
	return &Response{
		body:        body,
		Code:        502,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get establishment information
	// (GET /search)
	Search(w http.ResponseWriter, r *http.Request, params SearchParams) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// Search operation middleware
func (siw *ServerInterfaceWrapper) Search(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchParams

	// ------------- Required query parameter "filter" -------------

	if err := runtime.BindQueryParameter("form", true, true, "filter", r.URL.Query(), &params.Filter); err != nil {
		err = fmt.Errorf("invalid format for parameter filter: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "filter"})
		return
	}

	// ------------- Required query parameter "location" -------------

	if err := runtime.BindQueryParameter("form", true, true, "location", r.URL.Query(), &params.Location); err != nil {
		err = fmt.Errorf("invalid format for parameter location: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "location"})
		return
	}

	// ------------- Required query parameter "distance" -------------

	if err := runtime.BindQueryParameter("form", true, true, "distance", r.URL.Query(), &params.Distance); err != nil {
		err = fmt.Errorf("invalid format for parameter distance: %w", err)
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{err, "distance"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.Search(w, r, params)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter %s: %v", err.paramName, err.err)
}

func (err UnescapedCookieParamError) Unwrap() error { return err.err }

type UnmarshalingParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnmarshalingParamError) Error() string {
	return fmt.Sprintf("error unmarshaling parameter %s as JSON: %v", err.paramName, err.err)
}

func (err UnmarshalingParamError) Unwrap() error { return err.err }

type RequiredParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err RequiredParamError) Error() string {
	if err.err == nil {
		return fmt.Sprintf("query parameter %s is required, but not found", err.paramName)
	} else {
		return fmt.Sprintf("query parameter %s is required, but errored: %s", err.paramName, err.err)
	}
}

func (err RequiredParamError) Unwrap() error { return err.err }

type RequiredHeaderError struct {
	paramName string
}

// Error implements error.
func (err RequiredHeaderError) Error() string {
	return fmt.Sprintf("header parameter %s is required, but not found", err.paramName)
}

type InvalidParamFormatError struct {
	err       error
	paramName string
}

// Error implements error.
func (err InvalidParamFormatError) Error() string {
	return fmt.Sprintf("invalid format for parameter %s: %v", err.paramName, err.err)
}

func (err InvalidParamFormatError) Unwrap() error { return err.err }

type TooManyValuesForParamError struct {
	NumValues int
	paramName string
}

// Error implements error.
func (err TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("expected one value for %s, got %d", err.paramName, err.NumValues)
}

// ParameterName is an interface that is implemented by error types that are
// relevant to a specific parameter.
type ParameterError interface {
	error
	// ParamName is the name of the parameter that the error is referring to.
	ParamName() string
}

func (err UnescapedCookieParamError) ParamName() string  { return err.paramName }
func (err UnmarshalingParamError) ParamName() string     { return err.paramName }
func (err RequiredParamError) ParamName() string         { return err.paramName }
func (err RequiredHeaderError) ParamName() string        { return err.paramName }
func (err InvalidParamFormatError) ParamName() string    { return err.paramName }
func (err TooManyValuesForParamError) ParamName() string { return err.paramName }

type ServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type ServerOption func(*ServerOptions)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface, opts ...ServerOption) http.Handler {
	options := &ServerOptions{
		BaseURL:    "/",
		BaseRouter: chi.NewRouter(),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	for _, f := range opts {
		f(options)
	}

	r := options.BaseRouter
	wrapper := ServerInterfaceWrapper{
		Handler:          si,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Get("/search", wrapper.Search)
	})
	return r
}

func WithRouter(r chi.Router) ServerOption {
	return func(s *ServerOptions) {
		s.BaseRouter = r
	}
}

func WithServerBaseURL(url string) ServerOption {
	return func(s *ServerOptions) {
		s.BaseURL = url
	}
}

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request, err error)) ServerOption {
	return func(s *ServerOptions) {
		s.ErrorHandlerFunc = handler
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xU3WobTQx9lUHfB71ZvM5PoexdW0oaWtqSXAZfyLtyPOnOTyRtigl+96JZx2nipaVQ",
	"cuXxzpHOkXQ099CmkFOkqALNPUi7poDl+IE5sR0yp0ysnsrnQCJ4TXbUTSZoQJR9vIbttgKm28EzddBc",
	"7YGLbQXfemzpMBd2HZPIRK4KIgaavMjrpKnEeKUwHbz7gMy4sf+MajeP0DiEJfGB5kJa7XXtA/esi33u",
	"tLyhVuGhOPnsRZ+I+p9pBQ38Vz92uN61tx77cSDU5Pi4ShbekbTss/oUDZC90+SU8Y56l3uMkdgtsf1O",
	"sYMK1GtviZ4CoII7YhlTzGfz2ZFxpkwRs4cGTmZHsxMrDnVdNNdCyO3ajtdUyrFxoYk476CBy/HaIhgD",
	"KbFAc/Vc7Mek1EvlLkgUB8aoUrnE7q0qY2sga6036O1AvIGHacPK91pkPw5FeaBq58tJ0z1nf+91476M",
	"g5zi6FNbCvpblqlcnRfF2NJvc60SB1RDp2HZG/a5CRcWLjlFGffieD63nzZFpVjGgDn3ftRd34jVef8L",
	"wx+tNrqz+Otps75+Mkec/kO+8dmYoHqHnbug24GkbM3p0ZtDm5+/Cg6dEuZUQK9fQth5VOKIvbskviN2",
	"O6CxH79MW85Q6Ud5ACqQIQTkDTRwRupsg5a9l3WgqM7eBvOSRZZUUhSPOzhwDw3UttnbxfZnAAAA//+5",
	"WEXC2QUAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
