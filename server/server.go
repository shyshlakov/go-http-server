package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shyshlakov/go-http-server/config"
	"github.com/shyshlakov/go-http-server/handler"
)

type appServer struct {
	httpSrv *fiber.App
	h       *handler.HandlerRoutes
	cfg     *config.Config
}

func NewAppServer(h *handler.HandlerRoutes, cfg *config.Config) *appServer {
	httpSrv := fiber.New()
	httpSrv.Get("/api/tags", h.GetTags)
	httpSrv.Get("/api/articles/:slug", h.GetArticleBySlug)
	httpSrv.Get("/api/articles", h.GetArticles)

	httpSrv.Post("/api/articles", h.CreateArticle)
	httpSrv.Put("/api/articles/:slug", h.UpdateArticle)
	httpSrv.Delete("/api/articles/:slug", h.Delete)

	httpSrv.Post("/api/authors", h.CreateAuthor)
	httpSrv.Get("/api/authors/:name", h.GetAuthorByName)

	return &appServer{ //& - получить ссылку от объекта
		httpSrv: httpSrv,
		h:       h,
		cfg:     cfg,
	}
}

func (app *appServer) ListenAndServe() chan error {
	errCh := make(chan error)
	go func() {
		add := fmt.Sprintf(":%s", app.cfg.ServicePort)
		errCh <- app.httpSrv.Listen(add)
	}()
	return errCh
}
