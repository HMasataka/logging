package logging

import (
	"context"
	"sync"
)

type contextKey string

const loggingKey contextKey = "logging_key"

func WithValue(parent context.Context, key string, val any) context.Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	if v, ok := parent.Value(loggingKey).(*sync.Map); ok {
		mapCopy := copySyncMap(v)

		mapCopy.Store(key, val)

		return context.WithValue(parent, loggingKey, mapCopy)
	}

	v := &sync.Map{}
	v.Store(key, val)

	return context.WithValue(parent, loggingKey, v)
}

func copySyncMap(m *sync.Map) *sync.Map {
	var cp sync.Map

	m.Range(func(k, v any) bool {
		cp.Store(k, v)
		return true
	})

	return &cp
}
