package search

import (
	app "go-api-se-project/app"
	//log "github.com/sirupsen/logrus"
)

var srv *searchController

type searchController struct {
	cv *app.Configs
	em *app.ErrorMessage
	repo *searchRepo
}

func Init(cv *app.Configs, em *app.ErrorMessage) {
	srv = &searchController{
		cv: cv,
		em: em,
		repo: &searchRepo{cv: cv},
	}
} 

func NewInventoryService(cv *app.Configs, em *app.ErrorMessage) *searchController {
	repo := searchRepo{}

	return &searchController{
		cv: cv,
		em: em,
		repo: repo.InitSearchRepo(cv, em),
	}
}

func SearchLogic() (result searchModel, err error) {
	return
}