package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server struct {
	httpServer *http.Server
}

const defaultPort = ":8080"

func (s *Server) RunServer(handler http.Handler) error {

	port := viper.GetString("port")
	if port == "" {
		port = defaultPort
	}

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx *gin.Context) error {
	return s.httpServer.Shutdown(ctx)
}
