package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sfeir/storage/dto"
	"sfeir/storage/model"
)

type MovieStorage interface {
	AddMovie(
		ctx context.Context,
		movie dto.MovieDto,
	) (id primitive.ObjectID, err error)

	GetMovieByID(
		ctx context.Context,
		movieID primitive.ObjectID,
	) (movie model.Movie, err error)

	CountAllMovies(ctx context.Context) (movieCount int64, err error)

	GetMoviesPaginated(
		ctx context.Context,
		paginationDto dto.PaginationDto,
	) (movies *[]model.Movie, err error)
}
