package mock

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sfeir/global"
	"sfeir/storage/dto"
	"sfeir/storage/model"
)

type MoviesStorageImplMock struct{}

var (
	ValidMovieID, _ = primitive.ObjectIDFromHex("5f4d7c8b7c8b7c8b7c8b7c8b")
	MoviesCount     = 20
)

func NewMovieMock(movieID primitive.ObjectID) model.Movie {
	return model.Movie{
		ID:    movieID.Hex(),
		Title: "Movie_Title",
		Year:  2022,
		Director: &model.Director{
			FirstName: "Director_FirstName",
			LastName:  "Director_LastName",
		},
	}
}

func Movies() []model.Movie {
	movies := make([]model.Movie, MoviesCount)
	for i := 0; i < MoviesCount; i++ {
		movies = append(movies, NewMovieMock(primitive.NewObjectID()))
	}

	return movies
}

func (mStorage *MoviesStorageImplMock) GetMovieByID(
	ctx context.Context,
	movieID primitive.ObjectID,
) (movie model.Movie, err error) {

	if movieID != ValidMovieID {
		return model.Movie{}, global.InvalidIDErr(movieID.Hex())
	}

	return NewMovieMock(movieID), nil
}

func (mStorage *MoviesStorageImplMock) CountAllMovies(
	ctx context.Context,
) (movieCount int64, err error) {
	return int64(MoviesCount), nil
}

func (mStorage *MoviesStorageImplMock) GetMoviesPaginated(
	ctx context.Context,
	pagination dto.PaginationDto,
) (moviesResult *[]model.Movie, err error) {

	movies := Movies()

	moviesCount := int64(MoviesCount)

	startItemIndex := global.Min(pagination.Skip, moviesCount)
	lastItemIndex := global.Min(startItemIndex+pagination.Limit, moviesCount)

	moviesPaginated := movies[startItemIndex:lastItemIndex]

	return &moviesPaginated, nil
}

func (mStorage *MoviesStorageImplMock) AddMovie(
	ctx context.Context,
	movie dto.MovieDto,
) (id primitive.ObjectID, err error) {
	return ValidMovieID, nil
}
