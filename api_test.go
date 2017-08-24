package ctftime

import "testing"

func TestEventsAPIClient(t *testing.T) {
	var client interface{} = NewAPIClient("events", nil)

	if valid, ok := client.(APIClient); ok {
		valid.GetUrl()
		valid.GetAPIData()
	} else {
		t.Errorf("Invalid Typeof API Client %#v!", client)
	}
}
