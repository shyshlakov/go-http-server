package repo

import (
	"github.com/shyshlakov/go-http-server/persistence/model"
	"github.com/shyshlakov/go-http-server/restapi"
)

type Repo interface {
	Connect() error
	Close() error
	GetTags() ([]model.Tag, error)
	GetArticles() ([]model.Article, error)
	GetArticleBySlug(param string) (*model.Article, error)
	CreateArticle(params *restapi.ArticleRequestPayload) (*model.Article, error)
	UpdateArticle(slug string, params *restapi.ArticleRequestPayload) (*model.Article, error)
	DeleteArticle(param string) (bool, error)
	CreateAuthor(params *restapi.AuthorRequestPayload) (*model.Author, error)
	GetAuthorByName(param string) (*model.Author, error)
}
