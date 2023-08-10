package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/auth-api/pkg/e"
)

const (
	errServerStopped  = "Error on Server listen stop "
	errServerShutdown = "Error on Server shutdown"
)

type Server struct {
	srv *http.Server
}

func New(host, port string) *Server {
	srv := &http.Server{
		Addr:    host + ":" + port,
		Handler: router(),
	}
	return &Server{
		srv: srv,
	}
}

func (s *Server) Run() (err error) {
	defer func() { err = e.IsError(errServerStopped, err) }()

	err = s.srv.ListenAndServe()
	if err != http.ErrServerClosed {
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *Server) Shutdown() (err error) {
	defer func() { err = e.IsError(errServerShutdown, err) }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	err = s.srv.Shutdown(context.Background())
	if err != nil {

		return err
	}

	return nil
}
