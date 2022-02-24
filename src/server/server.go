package server

import (
	"fmt"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case PATH_API_CREATE_RECIVE:
			if isRequestGetApplicationJson(request) {
				responseOK(response, map[string]string{
					"hello": "World!",
				})
			} else {
				responseBadRequest(response, fmt.Errorf("Request isn't GET for application/json"))
			}
		default:
			http.NotFound(response, request)
		}
	})

	http.ListenAndServe(":8080", nil)
}
