/*!
 * Copyright 2019 - Author gauravm.git@gmail.com
 */

package web

import (
	"log"
	"net/http"

	"github.com/gauravmahto/go-start/utils"
	"github.com/gauravmahto/go-webserver/framework/router"
)

// Server Server entry point.
func Server() {

	log.Fatal(http.ListenAndServe(":8080", nil))

	utils.PrintLine("Server started")

	var httpMethod router.HTTPVerb

	httpMethod = "yes"

	utils.PrintLine("HTTP ", router.MethodGet)
	utils.PrintLine("HTTP ", router.MethodGet)
	utils.PrintLine("HTTP ", httpMethod)

}
