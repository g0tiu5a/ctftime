package ctftime

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
	return ""
}

func (client *top10APIClient) GetAPIData() interface{} {
	return nil
}
