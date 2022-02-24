package server

import "net/http"

func isRequestGetApplicationJson(req *http.Request) bool {
	return isRequestGet(req) && isRequestAcceptApplicationJson(req)
}

func isRequestGet(req *http.Request) bool {
	return req.Method == METHOD_GET
}

func isRequestAcceptApplicationJson(request *http.Request) bool {
	return request.Header.Get(ACCEPT) == APPLICATION_JSON
}

func isRequestApplicationJson(req *http.Request) bool {
	return req.Header.Get(CONTENT_TYPE) == APPLICATION_JSON
}
