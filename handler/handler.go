package handler

import (
	"fmt"

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
	res, err := h.s.GetTags()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (h *HandlerRoutes) GetArticles(ctx *fiber.Ctx) error {
	res, err := h.s.GetArticles()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot get articles: %v", err),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (h *HandlerRoutes) GetArticleBySlug(ctx *fiber.Ctx) error {
	param := ctx.Params("slug")
	if param == "idiot" {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Are you idiot?",
		})
	}
	res, err := h.s.GetArticleBySlug(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot get article by slug: %v", err),
		})
	}
	resPl := restapi.ResponsePayload{
		Data: res,
	}
	return ctx.Status(fiber.StatusOK).JSON(resPl)
}

func (h *HandlerRoutes) CreateArticle(ctx *fiber.Ctx) error {
	body := &restapi.ArticleRequestPayload{}
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot parse JSON: %v", err),
		})
	}
	res, err := h.s.CreateArticle(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot create article: %v", err),
		})
	}
	resPl := restapi.ResponsePayload{
		Data: res,
	}
	return ctx.Status(fiber.StatusCreated).JSON(resPl)
}

func (h *HandlerRoutes) UpdateArticle(ctx *fiber.Ctx) error {
	param := ctx.Params("slug")
	body := &restapi.ArticleRequestPayload{}
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

	res, err := h.s.UpdateArticle(param, body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot update article by slug: %v", err),
		})
	}
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

	res, err := h.s.DeleteArticle(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot delete article by slug: %v", err),
		})
	}
	return ctx.Status(fiber.StatusAccepted).JSON(res)
}

//Authors

func (h *HandlerRoutes) CreateAuthor(ctx *fiber.Ctx) error {
	body := &restapi.AuthorRequestPayload{}
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot parse JSON: %v", err),
		})
	}
	res, err := h.s.CreateAuthor(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot create author: %v", err),
		})
	}
	resPl := restapi.ResponsePayload{
		Data: res,
	}
	return ctx.Status(fiber.StatusCreated).JSON(resPl)
}

func (h *HandlerRoutes) GetAuthorByName(ctx *fiber.Ctx) error {
	param := ctx.Params("name")
	res, err := h.s.GetAuthorByName(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Cannot get author by name: %v", err),
		})
	}
	resPl := restapi.ResponsePayload{
		Data: res,
	}
	return ctx.Status(fiber.StatusOK).JSON(resPl)
}
