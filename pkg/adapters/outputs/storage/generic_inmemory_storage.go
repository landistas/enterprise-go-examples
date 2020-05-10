package storage

import "context"

type InfraStorage interface {
	Save(ctx context.Context, id string, v interface{}) error
	Read(ctx context.Context, id string) (interface{}, error)
	ListAll(ctx context.Context, filterFunc FilterFunc) (interface{}, error)
}

// FilterFunc check if the element must be filtered or not.
// It must return true when the element must be filtered and false when it must not be filtered
type FilterFunc func(element interface{}) bool

// NoFilterFunc returns always false so elements are not filtered
var NoFilterFunc = func(element interface{}) bool {
	return false
}
