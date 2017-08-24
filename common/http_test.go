package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDecodeJsonResponse(t *testing.T) {
	buf := GetTestData("event_1.json")

	// Create HTTP Response
	// 200 OK HTTP/1.0
	// <TEST_FILE json data>
	response := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       ioutil.NopCloser(bytes.NewReader(buf)),
	}

	var events []interface{}
	HttpResponseToStruct(response, &events)
}
