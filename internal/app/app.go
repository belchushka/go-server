package app

import (
	"fmt"
	"net/http"

	"lab18.matbea.xyz/savva/test-backend/internal/config"
	"lab18.matbea.xyz/savva/test-backend/internal/delivery/http/v1"
	"lab18.matbea.xyz/savva/test-backend/internal/delivery/http/server"
)

type app struct {
	name    string
	handler http.Handler
	config  *config.Config
}

type App interface {
	Run()
}

func NewApp() *app {

	config := config.GetConfig()
  handler := handler.NewHttpHandler(*&handler.NewHandlerParams{
    Cors: handler.Cors{
      AllowedMethods: config.Http.CORS.AllowedMethods,
      AllowedOrigins: config.Http.CORS.AllowedOrigins,
      AllowCredentials: config.Http.CORS.AllowCredentials,
      AllowedHeaders: config.Http.CORS.AllowedHeaders,
      Debug: config.Http.CORS.Debug,
    },
  })
  handler.RegisterRoutes()

	return &app{
		config: config,
    handler: handler.Engine,
	}
}

func (a *app) Run() {
	server := server.NewServer(&server.ServerConfig{
		ReadTimeout:  a.config.Http.ReadTimeout,
		WriteTimeout: a.config.Http.WriteTimeout,
		Port:         a.config.Http.Port,
		Host:         a.config.Http.Ip,
    Handler:      a.handler,
	})

  fmt.Printf("Starting http server at %s:%d \n", a.config.Http.Ip, a.config.Http.Port)

	server.Run()
}
