package svc

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sfeir/storage/dto"
	"sfeir/storage/model"
)

func (mSvc *movieSvc) AddMovie(
	movieDto dto.MovieDto,
) (movie model.Movie, err error) {

	movieRepo := mSvc.movieRepo

	insertedID, err := movieRepo.AddMovie(mSvc.ctx, movieDto)
	return movieRepo.GetMovieByID(mSvc.ctx, insertedID)
}

func (mSvc *movieSvc) GetMovieByID(
	movieID string,
) (movie model.Movie, err error) {
	movieRepo := mSvc.movieRepo

	movieObjectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return
	}

	return movieRepo.GetMovieByID(mSvc.ctx, movieObjectID)
}

func (mSvc *movieSvc) GetMovies(
	paginationDto dto.PaginationDto,
) (moviesPaginated model.MoviesPaginated, err error) {

	movieRepo := mSvc.movieRepo

	moviesCount, err := movieRepo.CountAllMovies(mSvc.ctx)
	if err != nil {
		return
	}
	movies, err := movieRepo.GetMoviesPaginated(mSvc.ctx, paginationDto)
	if err != nil {
		return
	}

	moviesPaginated = model.MoviesPaginated{
		TotalCount: moviesCount,
		Movies:     movies,
	}

	return
}
