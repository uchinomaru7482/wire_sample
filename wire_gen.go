// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"wiresample/handler"
	"wiresample/usecase"
)

// Injectors from wire.go:

func InitializeHelloHandler(ctx context.Context) handler.HelloHandler {
	helloUseCase := usecase.NewHelloUseCase(ctx)
	helloHandler := handler.NewHelloHandler(ctx, helloUseCase)
	return helloHandler
}