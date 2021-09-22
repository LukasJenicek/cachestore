package cachestore

import (
	"context"
	"errors"
)

var (
	MaxKeyLength = 80

	ErrKeyLengthTooLong = errors.New("cachestore: key length is too long")
	ErrInvalidKey       = errors.New("cachestore: invalid key")
)

type Store interface {
	// Set stores the given value associated to the key.
	Set(ctx context.Context, key string, value []byte) error
	// Get returns a stored value, or nil if the value is not assigned.
	Get(ctx context.Context, key string) ([]byte, error)
	// Returns true if the key exists.
	Exists(ctx context.Context, key string) (bool, error)
	// Delete removes a key and its associated value.
	Delete(ctx context.Context, key string) error

	// BatchSet sets all the values associated to the given keys.
	BatchSet(ctx context.Context, keys []string, values [][]byte) error
	// BatchGet returns the values of all the given keys at once. If any of the
	// keys has no value, nil is returned instead.
	BatchGet(ctx context.Context, keys []string) ([][]byte, error)
}
