package app

import (
	"github.com/riyan-eng/golang-boilerplate-one/internal/service"
)

type ServiceServer struct {
	exampleService service.ExampleService
	authService    service.AuthenticationService
}

func NewService(exampleService service.ExampleService, authService service.AuthenticationService) *ServiceServer {
	return &ServiceServer{
		exampleService: exampleService,
		authService:    authService,
	}
}
