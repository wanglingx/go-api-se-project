package search

import (
	//"github.com/jinzhu/gorm"
	app "go-api-se-project/app"
)

type searchRepo struct {
	cv *app.Configs
	em *app.ErrorMessage
}

func (repo *searchRepo) InitSearchRepo(cv *app.Configs, em*app.ErrorMessage) *searchRepo {
	return &searchRepo {
		cv: cv,
		em: em,
	}
}
	
func search(id int) (result searchModel, err error){
	if err = app.BetadineProject.DB.
		Table("farm").
		Select("*").
		Where("character_id = ?",id).
		Find(&result).
		Error; err != nil {
		return
	}
	return
}