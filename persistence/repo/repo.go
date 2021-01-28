package repo

import (
	"github.com/shyshlakov/go-http-server/persistence/model"
)

type Repo interface {
	Connect() error
	Close() error
	GetTags() ([]model.Tag, error)
	GetArticles() *[]model.Article
	GetArticleBySlug(param string) *model.Article
	CreateArticle(params *model.Article) *model.Article
	UpdateArticle(slug string, params *model.Article) *model.Article
	DeleteArticle(param string) bool
}
