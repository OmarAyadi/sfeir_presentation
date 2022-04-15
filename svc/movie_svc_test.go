package svc

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sfeir/global"
	"sfeir/storage/dto"
	"sfeir/storage/mock"
	"sfeir/storage/model"
	"sfeir/test"
	"testing"
)

func ShouldBeEqualMovies(t *testing.T, expectedMovie model.Movie, resultMovie model.Movie) {

	test.ShouldBeEqualOrFail(t, expectedMovie.Title, resultMovie.Title)
	test.ShouldBeEqualOrFail(t, expectedMovie.Year, resultMovie.Year)
	test.ShouldBeEqualOrFail(t, expectedMovie.Director, resultMovie.Director)
}

func TestAddMovie(t *testing.T) {
	mSvc := NewTestMovieSvc(t)

	insertNewECCaseDataTest := func(movieDto dto.MovieDto, expectedMovie model.Movie) {

		movie, err := mSvc.AddMovie(movieDto)

		test.ShouldBeNullOrFail(t, err)

		ShouldBeEqualMovies(t, expectedMovie, movie)
	}

	t.Run("insert new movie with no director", func(t *testing.T) {

		movieDto := dto.MovieDto{
			Title:       "Movie_Title",
			Year:        2022,
			DirectorDto: nil,
		}

		expectedMovie := model.Movie{
			Title:    "Movie_Title",
			Year:     2022,
			Director: nil,
		}

		insertNewECCaseDataTest(
			movieDto,
			expectedMovie,
		)
	})

	t.Run("insert new movie with a director", func(t *testing.T) {

		movieDto := dto.MovieDto{
			Title: "Movie_Title",
			Year:  2022,
			DirectorDto: &dto.DirectorDto{
				FirstName: "Director_FirstName",
				LastName:  "Director_LastName",
			},
		}

		expectedMovie := model.Movie{
			Title: "Movie_Title",
			Year:  2022,
			Director: &model.Director{
				FirstName: "Director_FirstName",
				LastName:  "Director_LastName",
			},
		}

		insertNewECCaseDataTest(
			movieDto,
			expectedMovie,
		)
	})
}

func TestGetMovies(t *testing.T) {

	mSvc := NewTestMovieSvc(t)

	for i := 0; i < 20; i++ {
		_, err := mSvc.AddMovie(dto.MovieDto{})
		test.ShouldBeNullOrFail(t, err)
	}

	getMoviesPaginated := func(paginationDto dto.PaginationDto, expectedMovieCount int) {

		movies, err := mSvc.GetMovies(paginationDto)
		test.ShouldBeNullOrFail(t, err)

		test.ShouldNotBeNullOrFail(t, movies.Movies)
		test.ShouldBeEqualOrFail(t, expectedMovieCount, len(*movies.Movies))
	}

	t.Run("should return 1 elements", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 1,
			Skip:  0,
		}

		getMoviesPaginated(paginationDto, 1)
	})

	t.Run("should return 10 elements", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 10,
			Skip:  0,
		}

		getMoviesPaginated(paginationDto, 10)
	})

	t.Run("should return 20 elements", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 20,
			Skip:  0,
		}

		getMoviesPaginated(paginationDto, 20)
	})

	t.Run("should return 20 elements even when the limit > number of elements in the database", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 200,
			Skip:  0,
		}

		getMoviesPaginated(paginationDto, 20)
	})

	t.Run("should return 10 elements after 10 skips", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 10,
			Skip:  10,
		}

		getMoviesPaginated(paginationDto, 10)
	})

	t.Run("should return 0 elements because we only have 20 elements in the database", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 10,
			Skip:  20,
		}

		getMoviesPaginated(paginationDto, 0)
	})

}

func TestGetMovieByID(t *testing.T) {

	mSvc := NewTestMovieSvc(t)

	t.Run("throw error because movie id is empty", func(t *testing.T) {

		_, err := mSvc.GetMovieByID(global.EmptyString)

		test.ShouldNotBeNullOrFail(t, err)
		test.ErrorShouldBeEqualOrFail(t, primitive.ErrInvalidHex, err)
	})

	t.Run("throw error because movie id not valid object id format", func(t *testing.T) {

		for _, invalidObjectIDFormat := range global.InvalidObjectIDsFormat {
			_, err := mSvc.GetMovieByID(invalidObjectIDFormat)

			test.ShouldNotBeNullOrFail(t, err)
			test.ErrorShouldBeEqualOrFail(t, primitive.ErrInvalidHex, err)
		}
	})

	t.Run("throw error because movie id not valid training case id", func(t *testing.T) {

		invalidMovieObjectID := primitive.NewObjectID()

		_, err := mSvc.GetMovieByID(invalidMovieObjectID.Hex())

		test.ShouldNotBeNullOrFail(t, err)
		test.ErrorShouldBeEqualOrFail(t, global.InvalidIDErr(invalidMovieObjectID.Hex()), err)
	})

	t.Run("insert new movie then fetch it by id", func(t *testing.T) {

		movieDto := dto.MovieDto{
			Title: "Movie_Title",
			Year:  2022,
			DirectorDto: &dto.DirectorDto{
				FirstName: "Director_FirstName",
				LastName:  "Director_LastName",
			},
		}

		movie, err := mSvc.AddMovie(movieDto)
		test.ShouldBeNullOrFail(t, err)

		movieAfterSearch, err := mSvc.GetMovieByID(movie.ID)
		test.ShouldBeNullOrFail(t, err)

		ShouldBeEqualMovies(t, movie, movieAfterSearch)
	})

}

func TestAddMovieWithMock(t *testing.T) {
	mSvc := NewMovieSvc(context.Background(), &mock.MoviesStorageImplMock{})

	movie, err := mSvc.AddMovie(dto.MovieDto{})

	test.ShouldBeNullOrFail(t, err)

	ShouldBeEqualMovies(t, mock.NewMovieMock(mock.ValidMovieID), movie)

}

func TestGetMoviesWithMock(t *testing.T) {
	mSvc := NewMovieSvc(context.Background(), &mock.MoviesStorageImplMock{})

	getMoviesPaginated := func(paginationDto dto.PaginationDto, expectedMovieCount int) {

		movies, err := mSvc.GetMovies(paginationDto)
		test.ShouldBeNullOrFail(t, err)

		test.ShouldNotBeNullOrFail(t, movies.Movies)
		test.ShouldBeEqualOrFail(t, expectedMovieCount, len(*movies.Movies))
	}

	t.Run("should return 1 elements", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 1,
			Skip:  0,
		}

		getMoviesPaginated(paginationDto, 1)
	})

	t.Run("should return 10 elements", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 10,
			Skip:  0,
		}

		getMoviesPaginated(paginationDto, 10)
	})

	t.Run("should return 20 elements", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 20,
			Skip:  0,
		}

		getMoviesPaginated(paginationDto, 20)
	})

	t.Run("should return 20 elements even when the limit > number of elements in the database", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 200,
			Skip:  0,
		}

		getMoviesPaginated(paginationDto, 20)
	})

	t.Run("should return 10 elements after 10 skips", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 10,
			Skip:  10,
		}

		getMoviesPaginated(paginationDto, 10)
	})

	t.Run("should return 0 elements because we only have 20 elements in the database", func(t *testing.T) {

		paginationDto := dto.PaginationDto{
			Limit: 10,
			Skip:  20,
		}

		getMoviesPaginated(paginationDto, 0)
	})
}

func TestGetMovieByIDWithMock(t *testing.T) {
	mSvc := NewMovieSvc(context.Background(), &mock.MoviesStorageImplMock{})

	t.Run("throw error because movie id is empty", func(t *testing.T) {

		_, err := mSvc.GetMovieByID(global.EmptyString)

		test.ShouldNotBeNullOrFail(t, err)
		test.ErrorShouldBeEqualOrFail(t, primitive.ErrInvalidHex, err)
	})

	t.Run("throw error because movie id not valid object id format", func(t *testing.T) {

		for _, invalidObjectIDFormat := range global.InvalidObjectIDsFormat {
			_, err := mSvc.GetMovieByID(invalidObjectIDFormat)

			test.ShouldNotBeNullOrFail(t, err)
			test.ErrorShouldBeEqualOrFail(t, primitive.ErrInvalidHex, err)
		}
	})

	t.Run("throw error because movie id not valid training case id", func(t *testing.T) {

		invalidMovieObjectID := primitive.NewObjectID()

		_, err := mSvc.GetMovieByID(invalidMovieObjectID.Hex())

		test.ShouldNotBeNullOrFail(t, err)
		test.ErrorShouldBeEqualOrFail(t, global.InvalidIDErr(invalidMovieObjectID.Hex()), err)
	})

	t.Run("insert new movie then fetch it by id", func(t *testing.T) {

		movieDto := dto.MovieDto{
			Title: "Movie_Title",
			Year:  2022,
			DirectorDto: &dto.DirectorDto{
				FirstName: "Director_FirstName",
				LastName:  "Director_LastName",
			},
		}

		movie, err := mSvc.AddMovie(movieDto)
		test.ShouldBeNullOrFail(t, err)

		movieAfterSearch, err := mSvc.GetMovieByID(movie.ID)
		test.ShouldBeNullOrFail(t, err)

		ShouldBeEqualMovies(t, movie, movieAfterSearch)
	})
}
