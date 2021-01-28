package postgres

import (
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
