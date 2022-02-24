package server

import "net/http"

const (
	METHOD_GET    = http.MethodGet
	METHOD_POST   = http.MethodPost
	METHOD_DELETE = http.MethodDelete

	STATUS_CODE_OK          = 200
	STATUS_CODE_NO_CONTENT  = 204
	STATUS_CODE_BAD_REQUEST = 400

	ACCEPT         = "Accept"
	CONTENT_TYPE   = "Content-Type"
	CONTENT_LENGTH = "Content-Length"

	APPLICATION_JSON = "application/json"
	TEXT_PLAIN       = "text/plain"

	PATH_API_CREATE_RECIVE = "/api/recipe"
)
