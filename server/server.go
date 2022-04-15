package server

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"sfeir/docs"
	"sfeir/global"
	"sfeir/server/rest"
	"sfeir/server/rest/handlers"
	"sfeir/storage"
	"sfeir/storage/mg"
	"strings"
	"time"
)

type Server struct {
	Engine       *gin.Engine
	Log          *logrus.Logger
	Storage      storage.Storage
	OidcProvider *oidc.Provider
}

func (s *Server) SetupStorage() {

	db, err := mg.OpenMongoConnectionAndGetDB(global.MongoHost, global.MongoDB)
	if err != nil {
		panic(errors.Wrap(err, "could not create mongo client"))
	}

	s.Storage = mg.NewStorageImpl(db)
}

func (s *Server) SetupEngin() {
	gin.SetMode(global.GinMode)
	s.Engine = gin.Default()

	docs.SwaggerInfo.BasePath = ""

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	s.SetupLog()
	s.SetupCors()
	s.SetupStorage()
}

func (s *Server) SetupLog() {
	s.Log = logrus.New()
	s.Log.SetLevel(logrus.InfoLevel)
}

func (s *Server) SetupCors() {
	s.Engine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "PUT", "PATCH", "HEAD", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"*",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		MaxAge: 90 * 24 * time.Hour,
	}))
}

func (s *Server) SetupRoute() {
	globalRouter := s.Engine.Group(global.EmptyString)
	rest.PublicRoutes(globalRouter)

	apiGroup := globalRouter.Group(global.ApiV1)

	handlers.MovieRoutes(apiGroup, s.Storage)
}

func NewServer() (s *Server, err error) {
	s = &Server{}
	s.SetupEngin()

	s.SetupRoute()

	return s, nil
}

func (s *Server) Run() (err error) {
	err = s.Engine.Run(":9009")
	return err
}
