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

func TestGetEventsData(t *testing.T) {
	client, err := newAPIClient("events", nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.GetAPIData()
	if err != nil {
		log.Fatal(err)
	}

	body, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
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
