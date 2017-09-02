package ctftime

import (
	"fmt"
	"log"
	"net/http"
)

type top10APIClient struct {
	Ctx APIContext
}

func newTop10APIClient(ctx APIContext) apiClient {
	return &top10APIClient{
		Ctx: ctx,
	}
}

func init() {
	registerAPIClient("top10", newTop10APIClient)
}

func (client *top10APIClient) GetUrl() string {
	url := fmt.Sprintf("%s/top/", API_ENDPOINT)
	if year, ok := client.Ctx["year"]; ok {
		url = fmt.Sprintf("%s%s/", url, year.(string))
	}

	return url
}

func (client *top10APIClient) GetAPIData() interface{} {
	url := client.GetUrl()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var top10s Top10s
	httpResponseToMap(resp, &top10s)
	if year, ok := client.Ctx["year"]; ok {
		var top10 Top10 = top10s[year.(string)]
		return top10
	} else {
		return top10s
	}
}
