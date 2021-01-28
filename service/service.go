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

func (s *AppService) GetTags() (*restapi.ResponsePayload, error) { //из service к repo
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

func (s *AppService) GetArticles() *[]model.Article {
	res := s.repo.GetArticles()
	return res
}

func (s *AppService) GetArticleBySlug(param string) *model.Article {
	res := s.repo.GetArticleBySlug(param)
	return res
}

func (s *AppService) CreateArticle(params *model.Article) *model.Article {
	return s.repo.CreateArticle(params)
}

func (s *AppService) UpdateArticle(slug string, params *model.Article) *model.Article {
	params.Slug = slug
	res := s.repo.UpdateArticle(slug, params)
	return res
}

func (s *AppService) DeleteArticle(param string) bool {
	res := s.repo.DeleteArticle(param)
	return res
}
