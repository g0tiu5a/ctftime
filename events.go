package ctftime

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	API_URL = API_ENDPOINT + "/events/"
	LIMIT   = 3
)

type EventsAPIClient struct{}

func (client *EventsAPIClient) GetUrl() string {
	now := time.Now().Unix()

	req, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("limit", fmt.Sprintf("%d", LIMIT))
	q.Add("start", strconv.FormatInt(now, 10))
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}

func (client *EventsAPIClient) GetAPIData() interface{} {
	url := client.GetUrl()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var events []Event
	HttpResponseToStruct(resp, &events)
	return events
}
