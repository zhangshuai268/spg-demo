package store

var client Factory

type Factory interface {
	Admin() AdminStore
	Code() CodeStore
	User() UserStore
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
