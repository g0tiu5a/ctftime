package ctftime

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestDecodeGetAPIData(t *testing.T) {
	result := GetAPIData()
	body, _ := json.Marshal(result)
	dummy_resp := &http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
	}
	var api_data []Event
	Decode(dummy_resp, &api_data)

	if len(api_data) != 3 {
		t.Error("Invalid event length!")
	}

	for _, event := range api_data {
		if reflect.TypeOf(Event{}) != reflect.TypeOf(event) {
			t.Errorf("Invalid event type of %v! (should be %v).", reflect.TypeOf(event), reflect.TypeOf(&Event{}))
		}
	}
}
