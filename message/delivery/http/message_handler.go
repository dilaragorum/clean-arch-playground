package http

import (
	"github.com/dilaragorum/clean-arch-playground/domain"
	"github.com/labstack/echo"
)

type MessageHandler struct {
	MUsecase domain.MessageUsecase
}

func NewMessageHandler(us domain.MessageUsecase) {
	handler := MessageHandler{MUsecase: us}
	_ = handler
}

func (m *MessageHandler) GetMessage(c echo.Context) error {
	return nil
}
