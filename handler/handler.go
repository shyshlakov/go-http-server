package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shyshlakov/go-http-server/persistence/model"
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
	res, err := h.s.GetTags()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (h *HandlerRoutes) GetArticles(ctx *fiber.Ctx) error {
	res := h.s.GetArticles()
	resPl := restapi.ResponsePayload{
		Data: res,
	}
	return ctx.Status(fiber.StatusOK).JSON(resPl)
}

func (h *HandlerRoutes) GetArticleBySlug(ctx *fiber.Ctx) error {
	param := ctx.Params("slug")
	if param == "idiot" {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Are you idiot?",
		})
	}
	res := h.s.GetArticleBySlug(param)
	resPl := restapi.ResponsePayload{
		Data: res,
	}
	return ctx.Status(fiber.StatusOK).JSON(resPl)
}

func (h *HandlerRoutes) CreateArticle(ctx *fiber.Ctx) error {
	body := &model.Article{}
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot parse JSON: %v", err),
		})
	}
	res := h.s.CreateArticle(body)
	resPl := restapi.ResponsePayload{
		Data: res,
	}
	return ctx.Status(fiber.StatusCreated).JSON(resPl)
}

func (h *HandlerRoutes) UpdateArticle(ctx *fiber.Ctx) error {
	param := ctx.Params("slug")
	body := &model.Article{}
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if param == "idiot" {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Are you idiot?",
		})
	}

	res := h.s.UpdateArticle(param, body)
	resPl := restapi.ResponsePayload{
		Data: res,
	}
	return ctx.Status(fiber.StatusAccepted).JSON(resPl)
}

func (h *HandlerRoutes) Delete(ctx *fiber.Ctx) error {
	param := ctx.Params("slug")
	if param == "idiot" {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Are you idiot?",
		})
	}

	h.s.DeleteArticle(param)
	return ctx.Status(fiber.StatusAccepted).JSON(nil)
}
