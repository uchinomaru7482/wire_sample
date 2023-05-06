package handler

import (
	"context"
	"log"

	"wiresample/usecase"
)

type HelloHandler interface {
	SayHello(ctx context.Context) error
}

type helloHandler struct {
	helloUseCase usecase.HelloUseCase
}

func NewHelloHandler(ctx context.Context, hu usecase.HelloUseCase) HelloHandler {
	return &helloHandler{
		helloUseCase: hu,
	}
}

func (hh helloHandler) SayHello(ctx context.Context) error {
	text, err := hh.helloUseCase.GetText(ctx)
	if err != nil {
		return nil
	}
	log.Print(text)
	return nil
}
