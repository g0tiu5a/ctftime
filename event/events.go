package event

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/g0tiu5a/ctftime/common"
)

const (
	LIMIT = 3
)

type EventsAPIClient struct{}

func (client *EventsAPIClient) GetUrl() string {
	now := time.Now().Unix()
	url := common.API_ENDPOINT + "/events/?limit=" + fmt.Sprintf("%d", LIMIT) + "&start=" + strconv.FormatInt(now, 10)
	return url
}

func (client *EventsAPIClient) GetAPIData() []Event {
	url := client.GetUrl()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var events []Event
	common.HttpResponseToStruct(response, &events)
	return events
}
