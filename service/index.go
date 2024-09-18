package service

type ServiceGroup struct {
	UserService UserService
	DXService   DXService
}

var ServiceGroupPtr = new(ServiceGroup)
