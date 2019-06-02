/*!
  * Copyright 2019 - Author gauravm.git@gmail.com
  */

package router

import (
	"net/http"
	"strings"

	"github.com/gauravmahto/go-start/utils"
)

// RegisterRouteArgs Arguments for RegisterRoute functions.
type RegisterRouteArgs struct {
	paths   []string
	handler http.Handler
}

// RegisterGet Register GET route.
func RegisterGet(args RegisterRouteArgs) {

	utils.PrintLine("RegisterGet()")

	http.Handle(strings.Join(args.paths, "/"), args.handler)

}
