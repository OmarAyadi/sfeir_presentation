package svc

import (
	"context"
	"sfeir/storage"
)

type movieSvc struct {
	ctx       context.Context
	movieRepo storage.MovieStorage
}

func NewMovieSvc(ctx context.Context, movieRepo storage.MovieStorage) movieSvc {
	return movieSvc{
		ctx:       ctx,
		movieRepo: movieRepo,
	}
}
