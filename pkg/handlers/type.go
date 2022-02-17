package handlers

import "context"

type ATMHandler struct{}

func NewATMHandler(ctx context.Context) ATMHandler {
	h := ATMHandler{}
	return h
}
