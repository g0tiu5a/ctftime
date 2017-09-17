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

func TestGetTop10Data(t *testing.T) {
	client, err := newAPIClient("top10", nil)
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

	var dummy_top10 Top10s
	httpResponseToMap(dummy_resp, &dummy_top10)
	if len(dummy_top10["2017"]) != 10 {
		t.Error("Invalid top10 of 2017 length!")
	}

	for _, team := range dummy_top10["2017"] {
		valid := reflect.TypeOf(Team{})
		actual := reflect.TypeOf(team)
		if actual != valid {
			t.Errorf("Invalid team type of %v! (should be %v).", actual, valid)
		}
	}
}
