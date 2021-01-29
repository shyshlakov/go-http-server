package service

import (
	"github.com/shyshlakov/go-http-server/persistence/model"
	"github.com/shyshlakov/go-http-server/persistence/repo"
	"github.com/shyshlakov/go-http-server/restapi"
)

type AppService struct {
	repo repo.Repo
}

func NewService(r repo.Repo) *AppService {
	return &AppService{
		repo: r,
	}
}

func (s *AppService) GetTags() (*restapi.ResponsePayload, error) {
	res, err := s.repo.GetTags()
	if err != nil {
		return nil, err
	}
	arr := []string{}
	for i := range res {
		arr = append(arr, res[i].Name)
	}
	return &restapi.ResponsePayload{
		Data: restapi.Tag{
			Tags: arr,
		},
	}, nil
}

func (s *AppService) GetArticles() (*restapi.ResponsePayload, error) {
	res, err := s.repo.GetArticles()
	if err != nil {
		return nil, err
	}
	return &restapi.ResponsePayload{
		Data: res,
	}, nil
}

func (s *AppService) GetArticleBySlug(param string) (*model.Article, error) {
	return s.repo.GetArticleBySlug(param)
}

func (s *AppService) CreateArticle(params *restapi.ArticleRequestPayload) (*model.Article, error) {
	return s.repo.CreateArticle(params)
}

func (s *AppService) UpdateArticle(slug string, params *restapi.ArticleRequestPayload) (*model.Article, error) {
	return s.repo.UpdateArticle(slug, params)
}

func (s *AppService) DeleteArticle(param string) (bool, error) {
	return s.repo.DeleteArticle(param)
}

func (s *AppService) CreateAuthor(params *restapi.AuthorRequestPayload) (*model.Author, error) {
	return s.repo.CreateAuthor(params)
}

func (s *AppService) GetAuthorByName(param string) (*model.Author, error) {
	return s.repo.GetAuthorByName(param)
}
