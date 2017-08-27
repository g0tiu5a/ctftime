package ctftime

import (
	"log"
)

type apiClient interface {
	GetUrl() string
	GetAPIData() interface{}
}

type APIContext map[string]interface{}
type apiClientFactory func(ctx APIContext) apiClient

var apiClientFactories = make(map[string]apiClientFactory)

func registerAPIClient(name string, factory apiClientFactory) {
	if factory == nil {
		log.Panicf("API Client Factory %s does not exist.", name)
	}

	_, registered := apiClientFactories[name]
	if registered {
		log.Panicf("API Client Factory %s already registered. Ignoring.", name)
	}

	apiClientFactories[name] = factory
}

func newAPIClient(name string, ctx APIContext) apiClient {
	clientFactory, ok := apiClientFactories[name]
	if !ok {
		log.Panicf("Invalid API Client name!")
	}

	return clientFactory(ctx)
}

func GetUrl(name string, ctx APIContext) string {
	client := newAPIClient(name, ctx)
	return client.GetUrl()
}

func GetAPIData(name string, ctx map[string]interface{}) interface{} {
	client := newAPIClient(name, ctx)
	return client.GetAPIData()
}
