package usecase

import (
	"context"
	"github.com/dilaragorum/clean-arch-playground/domain"
)

type messageUsecase struct{}

func NewMessageUsecase() domain.MessageUsecase {
	return &messageUsecase{}
}

func (m messageUsecase) GetMessage(ctx context.Context) (domain.Message, error) {
	//TODO implement me
	panic("implement me")
}
