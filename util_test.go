package ctftime

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"testing"
)

const (
	testFile = "event_1.json"
)

func TestGetTestData(t *testing.T) {
	buf := getTestData(testFile)

	data, err := ioutil.ReadFile(path.Join(test_dir, testFile))
	if err != nil {
		log.Fatal(err)
	}

	if string(buf) != string(data) {
		t.Errorf("Data doesn't match %v != %v\n", buf, data)
	}
}

func TestHttpResponseToStruct(t *testing.T) {
	buf := getTestData("event_1.json")

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
	httpResponseToStruct(response, &events)
}

func TestHttpResponseToMap(t *testing.T) {
	buf := getTestData("top10_1.json")

	response := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       ioutil.NopCloser(bytes.NewReader(buf)),
	}

	var top10 Top10
	httpResponseToMap(response, &top10)
}
