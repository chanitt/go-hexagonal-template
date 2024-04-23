package handler

import "github.com/chanitt/go-hexagonal-template/internal/core/port"

type handler struct {
	svc port.Service
}

func New(svc port.Service) Handler {
	return &handler{svc: svc}
}

type Handler interface{}