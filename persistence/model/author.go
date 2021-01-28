//nolint:unused,structcheck
package model

import (
	"time"

	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"
)

type Author struct {
	tableName  struct{}    `pg:"authors"`
	ID         uuid.UUID   `pg:"id,pk,type:uuid"`
	CreatedAt  time.Time   `pg:"created_at"`
	UpdatedAt  time.Time   `pg:"updated_at"`
	Name       string      `pg:"name"`
	RegisterOn pg.NullTime `pg:"register_on"`
	Image      string      `pg:"image"`
}

//body на создание из 3х последних полей
//структура на ответ из всех полей
