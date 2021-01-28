package postgres

import "github.com/shyshlakov/go-http-server/persistence/model"

func (r *postgreRepo) CreateArticle(params *model.Article) *model.Article {
	r.Article = append(r.Article, *params)
	return params
}

func (r *postgreRepo) UpdateArticle(slug string, params *model.Article) *model.Article {
	for i, v := range r.Article {
		if v.Slug == slug {
			r.Article[i] = *params
			return params
		}
	}
	return nil
}

func (r *postgreRepo) GetArticleBySlug(param string) *model.Article {
	for _, v := range r.Article {
		if v.Slug == param {
			return &v
		}
	}
	return nil
}

func (r *postgreRepo) DeleteArticle(param string) bool {
	for i, v := range r.Article {
		if v.Slug == param {
			r.Article = append(r.Article[:i], r.Article[i+1:]...)
			return true
		}
	}
	return false
}

func (r *postgreRepo) GetArticles() *[]model.Article {
	return &r.Article
}
