package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

var routeControllers = routes{
	route{
		"Index",
		"GET",
		"/",
		Index,
	},
	route{
		"OCR",
		"GET",
		"/ocr",
		OcrHandler,
	},
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, r := range routeControllers {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}

	return router
}
