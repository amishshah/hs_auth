// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unicsmcr/hs_auth/repositories"
	"github.com/unicsmcr/hs_auth/routers"
	"github.com/unicsmcr/hs_auth/routers/api/v1"
	"github.com/unicsmcr/hs_auth/services"
	"github.com/unicsmcr/hs_auth/utils"
)

// Injectors from wire.go:

func InitializeServer() (*gin.Engine, error) {
	logger, err := utils.NewLogger()
	if err != nil {
		return nil, err
	}
	database := utils.NewDatabase(logger)
	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)
	apiv1Router := v1.NewAPIV1Router(logger, userService)
	mainRouter := routers.NewMainRouter(logger, apiv1Router)
	engine := NewServer(mainRouter)
	return engine, nil
}
