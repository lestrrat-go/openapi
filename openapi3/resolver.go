package openapi3

import (
	"errors"
	"sync"
)

type Resolver struct {
	mu    sync.RWMutex
	cache map[string]interface{}
	ctx   OpenAPI
}

func NewResolver(ctx OpenAPI) *Resolver {
	return &Resolver{
		cache: make(map[string]interface{}),
		ctx:   ctx,
	}
}

func (r *Resolver) Resolve(path string) (interface{}, error) {
	r.mu.RLock()
	if v, ok := r.cache[path]; ok {
		defer r.mu.RUnlock()
		return v, nil
	}
	r.mu.RUnlock()

	v, ok := r.ctx.QueryJSON(path)
	if !ok {
		return nil, errors.New(`could not resolve reference`)
	}

	r.mu.Lock()
	r.cache[path] = v
	r.mu.Unlock()
	return v, nil
}
