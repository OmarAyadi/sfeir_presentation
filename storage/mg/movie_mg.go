package mg

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sfeir/global"
	"sfeir/storage/dto"
	"sfeir/storage/model"
)

func (mStorage *MoviesStorageImpl) GetMovieByID(
	ctx context.Context,
	movieID primitive.ObjectID,
) (movie model.Movie, err error) {

	filters := bson.D{
		{
			Key:   global.MongoID,
			Value: movieID,
		},
	}

	result := mStorage.MoviesCol.FindOne(ctx, filters)

	if result.Err() != nil {
		err = global.InvalidIDErr(movieID.Hex())
		return
	}

	err = result.Decode(&movie)

	return
}

func (mStorage *MoviesStorageImpl) CountAllMovies(
	ctx context.Context,
) (movieCount int64, err error) {
	return mStorage.MoviesCol.EstimatedDocumentCount(ctx)
}

func (mStorage *MoviesStorageImpl) GetMoviesPaginated(
	ctx context.Context,
	pagination dto.PaginationDto,
) (moviesResult *[]model.Movie, err error) {
	movies := make([]model.Movie, 0)
	moviesResult = &movies

	paginationOptions := &options.FindOptions{
		Limit: &pagination.Limit,
		Skip:  &pagination.Skip,
	}

	cursor, err := mStorage.MoviesCol.Find(ctx, bson.D{}, paginationOptions)
	if err != nil {
		return
	}

	for cursor.Next(ctx) {
		var movie model.Movie
		err = cursor.Decode(&movie)
		if err != nil {
			return
		}
		movies = append(movies, movie)
	}

	return
}

func (mStorage *MoviesStorageImpl) AddMovie(
	ctx context.Context,
	movie dto.MovieDto,
) (id primitive.ObjectID, err error) {

	insertResult, err := mStorage.MoviesCol.InsertOne(ctx, movie)
	if err != nil {
		return
	}

	id = insertResult.InsertedID.(primitive.ObjectID)
	return id, nil
}
