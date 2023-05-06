// +build wireinject

package main

import (
	"context"

	"wiresample/handler"
	"wiresample/usecase"

	"github.com/google/wire"
)

func InitializeHelloHandler(ctx context.Context) handler.HelloHandler {
    wire.Build(handler.NewHelloHandler, usecase.NewHelloUseCase)
    return nil
}
