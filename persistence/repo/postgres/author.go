package postgres

import (
	"time"

	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"
	"github.com/shyshlakov/go-http-server/persistence/model"
	"github.com/shyshlakov/go-http-server/restapi"
)

func (r *postgreRepo) CreateAuthor(params *restapi.AuthorRequestPayload) (*model.Author, error) {
	dst := &model.Author{}
	dst.ID = uuid.NewV4()
	dst.Name = params.Name
	t, err := time.Parse(time.RFC3339, params.RegisterOn)
	if err != nil {
		return nil, err
	}
	dst.RegisterOn = pg.NullTime{
		Time: t,
	}
	dst.Image = params.Image

	if _, err := r.db.Model(dst).Insert(); err != nil {
		return nil, err
	}
	return dst, nil
}

func (r *postgreRepo) GetAuthorByName(param string) (*model.Author, error) {
	dst := model.Author{}
	err := r.db.Model(&dst).Where("name = ?", param).Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return &dst, nil
}
