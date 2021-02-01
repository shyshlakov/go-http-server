package postgres

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shyshlakov/go-http-server/persistence/model"
	"github.com/shyshlakov/go-http-server/restapi"
)

func (r *postgreRepo) CreateArticle(params *restapi.ArticleRequestPayload) (*model.Article, error) {
	dst := &model.Article{}
	dst.ID = uuid.NewV4()
	dst.Slug = params.Slug
	dst.Title = params.Title
	dst.Description = params.Description
	dst.Body = params.Body
	dst.Favorited = params.Favorited
	dst.FavoritesCount = params.FavoritesCount
	//dst.Comment = params.Comment
	dst.Score = params.Score
	dst.LikedUsers = params.LikedUsers
	athId, err := uuid.FromString(params.AuthorId)
	if err != nil {
		return nil, err
	}
	dst.AuthorID = athId

	if _, err := r.db.Model(dst).Insert(); err != nil {
		return nil, err
	}
	if _, err := r.CreateTagInArticle(params.TagList, dst.ID); err != nil {
		return nil, err
	}
	res, err := r.GetArticleBySlug(params.Slug)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *postgreRepo) UpdateArticle(slug string, params *restapi.ArticleRequestPayload) (*model.Article, error) {
	dst, err := r.GetArticleBySlug(slug)
	if err != nil {
		return nil, err
	}
	dst.UpdatedAt = time.Now()
	dst.Slug = params.Slug
	dst.Title = params.Title
	dst.Description = params.Description
	dst.Body = params.Body
	dst.Favorited = params.Favorited
	dst.FavoritesCount = params.FavoritesCount
	//dst.Comment = params.Comment
	dst.Score = params.Score
	dst.LikedUsers = params.LikedUsers
	athId, err := uuid.FromString(params.AuthorId)
	if err != nil {
		return nil, err
	}
	dst.AuthorID = athId

	if _, err := r.db.Model(dst).Where("slug = ?", slug).Update(); err != nil {
		return nil, err
	}

	if _, err := r.DeleteTagInArticleByArticle(dst.ID); err != nil {
		return nil, err
	}
	if _, err := r.CreateTagInArticle(params.TagList, dst.ID); err != nil {
		return nil, err
	}
	res, err := r.GetArticleBySlug(params.Slug)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *postgreRepo) GetArticleBySlug(param string) (*model.Article, error) {
	dst := model.Article{}
	err := r.db.Model(&dst).Where("slug = ?", param).Limit(1).Relation("Author").Relation("TagList").Select()
	if err != nil {
		return nil, err
	}
	return &dst, nil
}

func (r *postgreRepo) DeleteArticle(param string) (bool, error) {
	dst := model.Article{}
	if _, err := r.db.Model(&dst).Where("slug = ?", param).Delete(); err != nil {
		return false, err
	}
	return true, nil
}

func (r *postgreRepo) GetArticles() ([]model.Article, error) {
	dst := []model.Article{}
	err := r.db.Model(&dst).Relation("Author").Relation("TagList").Select()
	if err != nil {
		return nil, err
	}
	return dst, nil
}
