package main

import (
	"myapp/pkg/server"

	"github.com/labstack/echo/v4"
)

func registerrouter(s *server.Server, appcontext *Appcontext) {
	v1 := s.Echo().Group("/v1/api")
	registerv1router(v1, appcontext)
}

func registerv1router(group *echo.Group, appcontext *Appcontext) {
	// reviewgroup:=group.Group("/review")
	// companygroup:=group.Group("/company")
	// usergroup:=group.Group("/user")

}
