//nolint:unused,structcheck
package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Tag struct {
	tableName struct{}  `pg:"tags"`
	ID        uuid.UUID `pg:"id,pk,type:uuid"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	Name      string    `pg:"name"`
}

type TagInArticle struct {
	tableName struct{}  `pg:"tag_in_articles"`
	ID        uuid.UUID `pg:"id,pk,type:uuid"`
	CreatedAt time.Time `pg:"created_at"`
	TagID     uuid.UUID `pg:"tag_id"`
	Tag       *Tag      `pg:"rel:has-one"`
	ArticleID uuid.UUID `pg:"article_id"`
	Article   *Article  `pg:"rel:has-one"`
}
