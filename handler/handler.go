package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shyshlakov/go-http-server/restapi"
	"github.com/shyshlakov/go-http-server/service"
)

type HandlerRoutes struct {
	s *service.AppService
}

func NewServerRoutes(s *service.AppService) *HandlerRoutes {
	return &HandlerRoutes{
		s: s,
	}
}

func (h *HandlerRoutes) GetTags(ctx *fiber.Ctx) error {
	res := h.s.GetTags()
	tags := restapi.Tags{
		Tags:     res,
		TagName:  "123123",
		TagId:    "9999",
		TestName: "Sasha",
	}
	return ctx.Status(fiber.StatusOK).JSON(tags)
}

func (h *HandlerRoutes) GetArticleBySlug(ctx *fiber.Ctx) error {
	param := ctx.Params("slug")
	return ctx.Status(fiber.StatusOK).JSON(param)
}
