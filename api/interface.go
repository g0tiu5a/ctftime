package main

type APIClient interface {
	GetUrl() string
	GetAPIData() interface{}
}
