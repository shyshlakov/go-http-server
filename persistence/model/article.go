//nolint:unused,structcheck
package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Article struct {
	tableName      struct{}               `pg:"articles"`
	ID             uuid.UUID              `pg:"id,pk,type:uuid"`
	CreatedAt      time.Time              `pg:"created_at"`
	UpdatedAt      time.Time              `pg:"updated_at"`
	Slug           string                 `pg:"slug"`
	Title          string                 `pg:"title"`
	Description    string                 `pg:"description"`
	Body           string                 `pg:"body"`
	TagList        []*Tag                 `pg:"many2many:tag_in_articles"`
	Favorited      bool                   `pg:"favorited"`
	FavoritesCount int                    `pg:"favorites_count"`
	Comment        map[string]interface{} `pg:"comment"`
	Score          float64                `pg:"score"`
	LikedUsers     []string               `pg:"liked_users"`
	AuthorID       uuid.UUID              `pg:"author_id"`
	Author         *Author                `pg:"rel:has-one"`
}
