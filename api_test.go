package ctftime

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestEventsAPIClient(t *testing.T) {
	client, err := newAPIClient("events", nil)
	if err != nil {
		log.Fatal(err)
	}

	if valid, ok := client.(apiClient); ok {
		valid.GetUrl()
		valid.GetAPIData()
	} else {
		t.Errorf("Invalid Typeof API Client %#v!", client)
	}
}

func TestGetUrl(t *testing.T) {
	apiUrl, err := GetUrl("events", nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = url.Parse(apiUrl)
	if err != nil {
		t.Errorf("Invalid URL %#v\n", apiUrl)
	}
}

func TestGetAPIData(t *testing.T) {
	data, err := GetAPIData("events", nil)
	if err != nil {
		log.Fatal(err)
	}

	body, err := json.Marshal(data)
	if err != nil {
		t.Errorf("Invalid JSON Format %#v\n", body)
	}

	dummy_resp := &http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
	}

	var dummy_events []Event
	httpResponseToStruct(dummy_resp, &dummy_events)
	if len(dummy_events) != 3 {
		t.Error("Invalid event length!")
	}

	for _, event := range dummy_events {
		valid := reflect.TypeOf(Event{})
		actual := reflect.TypeOf(event)
		if actual != valid {
			t.Errorf("Invalid event type of %v! (should be %v).", actual, valid)
		}
	}
}
