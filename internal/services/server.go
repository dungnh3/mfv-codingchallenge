package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dungnh3/mfv-codingchallenge/config"
	"github.com/dungnh3/mfv-codingchallenge/internal/repositories"
	l "github.com/dungnh3/mfv-codingchallenge/pkg/log"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type Server struct {
	server *http.Server
	r      repositories.Repository
	cfg    *config.Config
	logger l.Logger
}

func New(cfg *config.Config, r repositories.Repository) *Server {
	logger := l.New().Named("server")
	port := cfg.Server.HTTP.Port
	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.HandleMethodNotAllowed = true
	router.RemoveExtraSlash = true

	s := &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		r:      r,
		cfg:    cfg,
		logger: logger,
	}
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.GET("/health", s.healthCheck)
	router.GET("/live", s.liveCheck)

	usersRouter := router.Group("/users")
	{
		v1 := usersRouter.Group("/") // consider using api version in here -> /api/v1
		{
			v1.GET("/:id", s.healthCheck)
			v1.GET("/:id/accounts", s.healthCheck)
		}
	}

	accountsRouter := router.Group("/accounts")
	{
		{
			v1 := accountsRouter.Group("/") // consider using api version in here -> /api/v1
			{
				v1.GET("/:id", s.healthCheck)
			}
		}
	}
	return s
}

// Run application
func (s *Server) Run() error {
	s.logger.Info("Start the server at", l.Object("address", s.server.Addr))
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Close app and all the resources
func (s *Server) Close(ctx context.Context) error {
	s.logger.Info("Shutting down server...")
	return s.server.Shutdown(ctx)
}

func (s *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}

func (s *Server) liveCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
