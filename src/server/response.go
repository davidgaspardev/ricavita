package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func _responseWithJsonData(response http.ResponseWriter, payload []byte, statusCode int) {
	response.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	response.Header().Set(CONTENT_LENGTH, fmt.Sprint(len(payload)))
	if statusCode == STATUS_CODE_NO_CONTENT {
		statusCode = STATUS_CODE_OK
	}
	response.WriteHeader(statusCode)
	response.Write(payload)
}

func _responseWithErrorInfo(response http.ResponseWriter, err error, statusCode int) {
	errBuffer := []byte(err.Error())
	response.Header().Set(CONTENT_TYPE, TEXT_PLAIN)
	response.Header().Set(CONTENT_LENGTH, fmt.Sprint(len(errBuffer)))
	response.WriteHeader(statusCode)
	response.Write(errBuffer)
}

func responseBadRequest(response http.ResponseWriter, err error) {
	_responseWithErrorInfo(response, err, http.StatusBadRequest)
}

func responseUnauthorized(response http.ResponseWriter) {
	_responseWithErrorInfo(response, fmt.Errorf("You don't have permission"), http.StatusUnauthorized)
}

func responseInternalServerError(response http.ResponseWriter, err error) {
	_responseWithErrorInfo(response, err, http.StatusInternalServerError)
}

func responseOK(response http.ResponseWriter, data interface{}) {
	payload, err := dataToPayload(data)
	if err != nil {
		responseInternalServerError(response, err)
	} else {
		_responseWithJsonData(response, payload, STATUS_CODE_OK)
	}
}

func dataToPayload(data interface{}) (payload []byte, err error) {
	switch fmt.Sprintf("%T", data) {
	case "string":
		payload = []byte(data.(string))
	case "[]uint8":
		payload = data.([]byte)
	default:
		payload, err = json.Marshal(data)
	}
	return
}
