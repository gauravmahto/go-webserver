/*!
  * Copyright 2019 - Author gauravm.git@gmail.com
  */

package web

import (
	"log"
	"net/http"

	"github.com/gauravmahto/go-webserver/router"
	"github.com/gauravmahto/go-start/utils"
)

// Server Server entry point.
func Server() {

	log.Fatal(http.ListenAndServe(":8080", nil))

	utils.PrintLine("Server started")

}
