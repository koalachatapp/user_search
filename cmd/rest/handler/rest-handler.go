package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koalachatapp/usersearch/internal/core/port"
)

type RestHandler struct {
	service port.UsersearchService
}

func NewRestHandler(service port.UsersearchService) *RestHandler {
	return &RestHandler{
		service: service,
	}
}

func (r *RestHandler) Get(ctx *fiber.Ctx) error {
	return nil
}
