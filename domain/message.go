package domain

import "context"

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type MessageUsecase interface {
	GetMessage(ctx context.Context) (Message, error)
}
