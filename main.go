package main

import (
	"context"

	// "wiresample/handler"
	// "wiresample/usecase"
)

func main() {
	ctx := context.Background()
	// hh, _ := initialize(ctx)

	hh := InitializeHelloHandler(ctx)
	hh.SayHello(ctx)
}

// func initialize(ctx context.Context) (handler.HelloHandler, error) {
// 	hu := usecase.NewHelloUseCase(ctx)
// 	hh := handler.NewHelloHandler(ctx, hu)
// 	return hh, nil
// }
