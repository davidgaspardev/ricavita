package server

import (
	"fmt"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case PATH_API_CREATE_RECIVE:
			routePathApiCreateRecive(response, request)
		default:
			http.NotFound(response, request)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func routePathApiCreateRecive(response http.ResponseWriter, request *http.Request) {
	if !isRequestGetApplicationJson(request) {
		responseBadRequest(response, fmt.Errorf("Request isn't GET for application/json"))
		return
	}

	responseOK(response, map[string]string{
		"hello": "World!",
	})
}
