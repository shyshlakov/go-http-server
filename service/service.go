package service

import (
	"github.com/shyshlakov/go-http-server/persistence/repo"
)

type AppService struct {
	Repo *repo.PostgreRepo
}

func (s *AppService) GetTags() []string { //из service к repo
	// res := make([]string, 0)
	// res = append(res, "123")
	res := []string{
		"reactjs",
		"angularjs",
	}
	return res
}
