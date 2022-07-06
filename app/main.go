package main

import (
	_messageHttpDelivery "github.com/dilaragorum/clean-arch-playground/message/delivery/http"
	_messageUcase "github.com/dilaragorum/clean-arch-playground/message/usecase"

	"github.com/labstack/echo"
)

func main() {
	mu := _messageUcase.NewMessageUsecase()
	_messageHttpDelivery.NewMessageHandler(mu)

	e := echo.New()
	_ = e

}
