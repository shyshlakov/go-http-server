package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shyshlakov/go-http-server/handler"
)

type appServer struct {
	httpSrv *fiber.App
	h       *handler.HandlerRoutes
}

func NewAppServer(h *handler.HandlerRoutes) *appServer {
	httpSrv := fiber.New()
	httpSrv.Get("/api/tags", h.GetTags)
	httpSrv.Get("/api/articles/:slug", h.GetArticleBySlug)
	return &appServer{ //& - получить ссылку от объекта
		httpSrv: httpSrv,
		h:       h,
	}
}

func (app *appServer) ListenAndServe() chan error {
	errCh := make(chan error)
	go func() {
		errCh <- app.httpSrv.Listen(":8082")
	}()
	return errCh
}
