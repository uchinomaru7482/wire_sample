package usecase

import (
	"context"
)

type HelloUseCase interface {
	GetText(ctx context.Context) (string, error)
}

type helloUseCase struct {}

func NewHelloUseCase(ctx context.Context) HelloUseCase {
	return &helloUseCase{}
}

func (mu helloUseCase) GetText(ctx context.Context) (string, error) {
	return "hello, uchinomaru !!!!!", nil
}
