package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sfeir/global"
	"sfeir/storage"
	"sfeir/storage/dto"
	"sfeir/svc"
)

func MovieRoutes(
	apiGroup *gin.RouterGroup,
	movieStorage storage.MovieStorage,
) {

	pathByID := fmt.Sprintf("/:%s", global.MovieIDParam)
	movieGroup := apiGroup.Group("/movies")
	{
		movieGroup.POST("", AddMovie(movieStorage))
		movieGroup.GET("", GetMovies(movieStorage))
		movieGroup.GET(pathByID, GetMovieByID(movieStorage))
	}
}

// AddMovie godoc
// @Summary Insert a new movie
// @Description Insert a new movie
// @Tags Movies
// @Accept json
// @Produce json
// @Param request body dto.MovieDto true "body"
// @Success 200 {object} model.Movie
// @Failure 400 {object} dto.ServerError
// @Failure 500 {object} dto.ServerError
// @Router /api/v1/movies [post]
func AddMovie(movieStorage storage.MovieStorage) gin.HandlerFunc {
	return func(c *gin.Context) {

		movieDto := dto.MovieDto{}
		err := c.ShouldBindJSON(&movieDto)
		if err != nil {
			c.AbortWithStatusJSON(400, dto.ErrorReturn(err))
			return
		}

		movieSvc := svc.NewMovieSvc(
			c.Request.Context(),
			movieStorage,
		)

		movie, err := movieSvc.AddMovie(movieDto)

		if err != nil {
			c.AbortWithStatusJSON(400, dto.ErrorReturn(err))
			return
		}

		c.JSON(200, movie)
	}
}

// GetMovies godoc
// @Summary Get all Movies paginated
// @Description Get all Movies paginated
// @Tags Movies
// @Accept json
// @Produce json
// @Param page query int64 false "index of the page" default(0)
// @Param limit query int64 false "number of elements to return in a page" default(20)
// @Success 200 {object} model.MoviesPaginated
// @Failure 400 {object} dto.ServerError
// @Failure 500 {object} dto.ServerError
// @Router /api/v1/movies [get]
func GetMovies(movieStorage storage.MovieStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieSvc := svc.NewMovieSvc(
			c.Request.Context(),
			movieStorage,
		)

		paginationDto := dto.NewPaginationDto(c)

		evergreenCases, err := movieSvc.GetMovies(paginationDto)
		if err != nil {
			c.AbortWithStatusJSON(400, dto.ErrorReturn(err))
			return
		}

		c.JSON(200, evergreenCases)
	}
}

// GetMovieByID godoc
// @Summary Get Movie by ID
// @Description Get Movie by ID
// @Tags Movies
// @Accept json
// @Produce json
// @Param movie_id path string true "movie_id"
// @Success 200 {object} model.Movie
// @Failure 400 {object} dto.ServerError
// @Failure 500 {object} dto.ServerError
// @Router /api/v1/movies/{movie_id} [get]
func GetMovieByID(movieStorage storage.MovieStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieID := c.Param(global.MovieIDParam)

		movieSvc := svc.NewMovieSvc(
			c.Request.Context(),
			movieStorage,
		)

		evergreenCase, err := movieSvc.GetMovieByID(
			movieID,
		)
		if err != nil {
			c.AbortWithStatusJSON(400, dto.ErrorReturn(err))
			return
		}

		c.JSON(200, evergreenCase)
	}
}
