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

// RegisterRouteArgs Arguments for RegisterRoute functions.
type RegisterRouteArgs struct {
	paths   []string
	handler http.HandlerFunc
}

// RegisterGet Register GET route.
func RegisterGet(args RegisterRouteArgs) {

	utils.PrintLine("RegisterGet()")

	handler := func(responseW http.ResponseWriter, request *http.Request) {

		if http.MethodGet == request.Method {

			args.handler(responseW, request)

		} else {

			utils.PrintLine("Invalid HTTP verb")

		}

	}

	http.HandleFunc(strings.Join(args.paths, "/"), handler)

}
