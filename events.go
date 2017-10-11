package ctftime

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type eventsAPIClient struct {
	Ctx APIContext
}

func newEventsAPIClient(ctx APIContext) apiClient {
	return &eventsAPIClient{
		Ctx: ctx,
	}
}

func init() {
	registerAPIClient("events", newEventsAPIClient)
}

func (client *eventsAPIClient) GetUrl() (string, error) {
	if event_id, ok := client.Ctx["event_id"]; ok {
		// 特定のイベントを取得するAPI
		// https://ctftime.org/api/v1/events/{event_id}/
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("%s/events/%d/", API_ENDPOINT, event_id.(int)),
			nil,
		)
		if err != nil {
			return "", err
		}
		return req.URL.String(), nil
	} else {
		// 特定期間の複数のイベントを取得するAPI
		// https://ctftime.org/api/v1/events/?limit={number}&start={timestamp}&finish={timestamp}
		now := time.Now().Unix()

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/events/", API_ENDPOINT), nil)
		if err != nil {
			return "", err
		}

		q := req.URL.Query()
		q.Add("limit", fmt.Sprintf("%d", LIMIT))
		q.Add("start", strconv.FormatInt(now, 10))
		req.URL.RawQuery = q.Encode()
		return req.URL.String(), nil
	}
}

func (client *eventsAPIClient) GetAPIData() (interface{}, error) {
	url, err := client.GetUrl()
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if _, ok := client.Ctx["event_id"]; ok {
		var event Event
		httpResponseToStruct(resp, &event)
		return event, nil
	} else {
		var events Events
		httpResponseToStruct(resp, &events)
		return events, nil
	}
}
