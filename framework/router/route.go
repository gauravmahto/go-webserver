/*!
 * Copyright 2019 - Author gauravm.git@gmail.com
 */

package router

import (
	"net/http"
	"strings"

	"github.com/gauravmahto/go-start/utils"
)

// Router The Router.
type Router struct {
}

// HTTPVerb Various HTTP methods
type HTTPVerb string

// HTTP methods
const (
	MethodGet     HTTPVerb = "GET"
	MethodHead    HTTPVerb = "HEAD"
	MethodPost    HTTPVerb = "POST"
	MethodPut     HTTPVerb = "PUT"
	MethodPatch   HTTPVerb = "PATCH"
	MethodDelete  HTTPVerb = "DELETE"
	MethodConnect HTTPVerb = "CONNECT"
	MethodOptions HTTPVerb = "OPTIONS"
	MethodTrace   HTTPVerb = "TRACE"
)

// RegisterRouteArgs Arguments for RegisterRoute functions.
type RegisterRouteArgs struct {
	verb    HTTPVerb
	paths   []string
	handler http.HandlerFunc
}

// RegisterGet Register GET route.
func RegisterGet(args RegisterRouteArgs) {

	utils.PrintLine("RegisterGet()")

	handler := func(responseW http.ResponseWriter, request *http.Request) {

		if args.verb == HTTPVerb(request.Method) {

			args.handler(responseW, request)

		} else {

			utils.PrintLine("Invalid HTTP verb")

		}

	}

	http.HandleFunc(strings.Join(args.paths, "/"), handler)

}
