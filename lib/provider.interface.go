package lib

type OpenDbConfig struct {
	url      *string
	host     *string
	username *string
	port     *int
	password *string
}

type ProviderInterface interface {
	OpenDb(config OpenDbConfig)
	RunQuery(query string)
}
