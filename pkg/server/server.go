package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

// type config struct {
// 	address string
// 	port int32
// }

type Server struct {
	echo    *echo.Echo
	addr    string //its addr:port 127.0.0.1:8000
	healthy bool
}

func (s Server) Start() error {
	s.healthy = true
	return s.echo.Start(s.addr)
}

func (s Server) Stop(ctx context.Context) error {
	s.healthy = false
	return s.echo.Shutdown(ctx)
}

func (s *Server) Echo() *echo.Echo {
	return s.echo
}

func NewServer(address string, port int32) (*Server, error) {
	e := echo.New()
	addr := fmt.Sprintf("%s:%d", address, port)
	return &Server{
		echo: e,
		addr: addr,
	}, nil
}
