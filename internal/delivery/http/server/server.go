package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	httpServer *http.Server
  config *ServerConfig
}

type Server interface {
	Run()
}

type ServerConfig struct {
  Port int
  Host string
  ReadTimeout int
  WriteTimeout int
  Handler http.Handler
}

func NewServer(config *ServerConfig) *server {

	httpServer := &http.Server{
		MaxHeaderBytes: 1 << 20,
    Handler: config.Handler,
    ReadTimeout: time.Second * time.Duration(config.ReadTimeout),
    WriteTimeout: time.Second * time.Duration(config.WriteTimeout),
	}

	return &server{
		httpServer: httpServer,
    config: config,
	}
}

func (s *server) Run() {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.config.Host, fmt.Sprint(s.config.Port)))

  if err != nil {
    panic(fmt.Sprintf("Failed to start server due to error: %s", err))
  }

	errChan := make(chan error)

	go func() {
		if err := s.httpServer.Serve(listener); err != nil {
			errChan <- err
		}
	}()

	exitChan := make(chan os.Signal)

	signal.Notify(exitChan, syscall.SIGKILL, syscall.SIGTERM, os.Interrupt)

	select {
	case <-exitChan:
		fmt.Println("Stopping server due to system call")
	case err := <-errChan:
		fmt.Printf("Stopping server due to error: %s", err)
	}

	s.httpServer.Shutdown(context.Background())
}
