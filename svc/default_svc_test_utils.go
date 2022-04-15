package svc

import (
	"context"
	"sfeir/storage/mg"
	"testing"
)

func NewTestMovieSvc(t *testing.T) movieSvc {
	t.Helper()

	db := mg.OpenTestDataBaseHelper(t)
	movieStorage := mg.NewMoviesStorageImpl(db)

	return NewMovieSvc(context.Background(), &movieStorage)
}
