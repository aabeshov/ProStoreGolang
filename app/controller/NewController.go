package Controller

import "GolangwithFrame/src/app/service"

type Controller struct {
	service service.Service
}

func New(service service.Service) Controller {
	return Controller{
		service: service,
	}
}
