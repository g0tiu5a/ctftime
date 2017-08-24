package ctftime

import (
	"log"
	"sync"
)

type APIClient interface {
	GetUrl() string
	GetAPIData() interface{}
}

type APIClientFactory func(ctx map[string]interface{}) APIClient

var (
	APIClientFactories = make(map[string]APIClientFactory)
	once               sync.Once
)

func registerAPIClient(name string, factory APIClientFactory) {
	if factory == nil {
		log.Panicf("API Client Factory %s does not exist.", name)
	}

	_, registered := APIClientFactories[name]
	if registered {
		log.Panicf("API Client Factory %s already registered. Ignoring.", name)
	}

	APIClientFactories[name] = factory
}

func NewAPIClient(name string, ctx map[string]interface{}) APIClient {
	once.Do(func() {
		registerAPIClient("events", newEventsAPIClient)
	})

	clientFactory, ok := APIClientFactories[name]
	if !ok {
		log.Panicf("Invalid API Client name!")
	}

	return clientFactory(ctx)
}
