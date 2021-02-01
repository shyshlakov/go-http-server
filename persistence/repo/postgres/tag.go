package postgres

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shyshlakov/go-http-server/persistence/model"
)

func (r *postgreRepo) GetTags() ([]model.Tag, error) {
	// number := rand.Intn(10)
	// slice := make([]string, 0)
	// for i := 0; i < number; i++ {
	// 	slice = append(slice, "wow")
	// }
	// return nil, slice
	dst := []model.Tag{}
	err := r.db.Model(&dst).Select()
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func (r *postgreRepo) CreateTagInArticle(tags []string, article uuid.UUID) ([]string, error) {
	for _, v := range tags {
		dst := &model.TagInArticle{}
		dst.ID = uuid.NewV4()
		tagId, tagErr := uuid.FromString(v)
		if tagErr != nil {
			return nil, tagErr
		}
		dst.TagID = tagId
		dst.ArticleID = article
		if _, err := r.db.Model(dst).Insert(); err != nil {
			return nil, err
		}
	}
	return tags, nil
}

func (r *postgreRepo) DeleteTagInArticleByArticle(article uuid.UUID) (uuid.UUID, error) {
	dst := model.TagInArticle{}
	if _, err := r.db.Model(&dst).Where("article_id = ?", article).Delete(); err != nil {
		return article, err
	}
	return article, nil
}
