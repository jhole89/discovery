package main

import (
	"net/http"
)

func registerRoutes() {

	http.HandleFunc("/query", dynamicQuery)
}
