package ctftime

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestAPIClient(t *testing.T) {
	var client interface{} = &EventsAPIClient{}

	if valid, ok := client.(APIClient); ok {
		valid.GetUrl()
		valid.GetAPIData()
	} else {
		t.Errorf("Invalid Typeof API Client %#v!", client)
	}
}

func TestGetEventsData(t *testing.T) {
	client := &EventsAPIClient{}

	result := client.GetAPIData()

	body, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	resp := &http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
	}

	var events []Event
	HttpResponseToStruct(resp, &events)
	if len(events) != 3 {
		t.Error("Invalid event length!")
	}

	for _, event := range events {
		valid := reflect.TypeOf(Event{})
		actual := reflect.TypeOf(event)
		if actual != valid {
			t.Errorf("Invalid event type of %v! (should be %v).", actual, valid)
		}
	}
}
