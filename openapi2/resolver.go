package openapi2

import (
	"errors"
	"sync"
)

type resolver struct {
	mu    sync.RWMutex
	cache map[string]interface{}
	ctx   Swagger
}

func NewResolver(ctx Swagger) Resolver {
	return &resolver{
		cache: make(map[string]interface{}),
		ctx:   ctx,
	}
}

func (r *resolver) Resolve(path string) (interface{}, error) {
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
