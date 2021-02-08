package server

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/shyshlakov/go-http-server/config"
	"github.com/shyshlakov/go-http-server/handler"
	"github.com/shyshlakov/go-http-server/middleware"
	"github.com/shyshlakov/go-http-server/restapi"
)

type appServer struct {
	httpSrv *fiber.App
	h       *handler.HandlerRoutes
	cfg     *config.Config
}

func NewAppServer(h *handler.HandlerRoutes, cfg *config.Config, swagger *openapi3.Swagger) *appServer {
	httpSrv := fiber.New()
	swagger.Servers = nil
	httpSrv.Use(middleware.OapiRequestValidator(swagger))

	// httpSrv.Get("/api/tags", h.GetTags)
	// httpSrv.Get("/api/articles/:slug", h.GetArticleBySlug)
	// httpSrv.Get("/api/articles", h.GetArticles)

	// httpSrv.Post("/api/articles", h.CreateArticle)
	// httpSrv.Put("/api/articles/:slug", h.UpdateArticle)
	// httpSrv.Delete("/api/articles/:slug", h.Delete)

	// httpSrv.Post("/api/authors", h.CreateAuthor)
	// httpSrv.Get("/api/authors/:name", h.GetAuthorByName)

	restapi.RegisterHandlers(httpSrv, h)
	return &appServer{
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
