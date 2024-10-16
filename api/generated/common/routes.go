// Package common provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package common

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns 200 if healthy.
	// (GET /health)
	MakeHealthCheck(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// MakeHealthCheck converts echo context to params.
func (w *ServerInterfaceWrapper) MakeHealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.MakeHealthCheck(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/health", wrapper.MakeHealthCheck, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
