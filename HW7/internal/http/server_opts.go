package http

import (
	"HW7/internal/cache"
	"HW7/internal/store"
)

type ServerOption func(srv *Server)

func WithAddress(address string) ServerOption {
	return func(srv *Server) {
		srv.Address = address
	}
}

func WithStore(store store.FactsRepository) ServerOption {
	return func(srv *Server) {
		srv.store = store
	}
}

func WithCache(cache cache.Cache) ServerOption {
	return func(srv *Server) {
		srv.cache = cache
	}
}