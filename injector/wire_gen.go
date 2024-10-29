// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"mowa-backend/api/controller"
	"mowa-backend/api/services"
)

// Injectors from injector.go:

func InitializeUserController() controller.UserController {
	userService := services.NewUserService()
	userController := controller.NewUserController(userService)
	return userController
}