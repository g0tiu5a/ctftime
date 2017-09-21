package ctftime

import (
	"errors"
	"log"
)

type apiClient interface {
	GetUrl() (string, error)
	GetAPIData() (interface{}, error)
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

func newAPIClient(name string, ctx APIContext) (apiClient, error) {
	clientFactory, ok := apiClientFactories[name]
	if !ok {
		return nil, errors.New("Invalid API Client name!")
	}

	return clientFactory(ctx), nil
}

func GetUrl(name string, ctx APIContext) (string, error) {
	client, err := newAPIClient(name, ctx)
	if err != nil {
		return "", err
	}
	url, err := client.GetUrl()
	if err != nil {
		return "", err
	}
	return url, nil
}

func GetAPIData(name string, ctx APIContext) (interface{}, error) {
	client, err := newAPIClient(name, ctx)
	if err != nil {
		return nil, err
	}
	data, err := client.GetAPIData()
	if err != nil {
		return nil, err
	}
	return data, nil
}
