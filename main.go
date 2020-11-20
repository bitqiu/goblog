package main

import (
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"
)

var router *mux.Router

func main() {

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}