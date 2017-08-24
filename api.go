package ctftime

type APIClient interface {
	GetUrl() string
	GetAPIData() interface{}
}
