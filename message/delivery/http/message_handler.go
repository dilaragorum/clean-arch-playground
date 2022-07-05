package http

import "github.com/dilaragorum/clean-arch-playground/domain"

type MessageHandler struct {
	MUsecase domain.MessageUsecase
}

func NewMessageHandler(us domain.MessageUsecase) {
	handler := MessageHandler{MUsecase: us}
	_ = handler
}
