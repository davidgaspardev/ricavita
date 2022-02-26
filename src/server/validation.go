package server

import "net/http"

func isRequestGetApplicationJson(req *http.Request) bool {
	return isRequestGet(req) && isRequestAcceptApplicationJson(req)
}

func isRequestGet(req *http.Request) bool {
	return req.Method == METHOD_GET
}

func isRequestAcceptApplicationJson(request *http.Request) bool {
	accept := request.Header.Get(ACCEPT)
	return accept == APPLICATION_JSON || accept == "" || accept == "*/*"
}

func isRequestApplicationJson(req *http.Request) bool {
	return req.Header.Get(CONTENT_TYPE) == APPLICATION_JSON
}
