package repositories

import (
	"errors"
	"github.com/OahcUil94/go-boilerplate/datamodels"
	"github.com/OahcUil94/go-boilerplate/datasource"
	"github.com/jinzhu/gorm"
)

type IHeroRepository interface {
	SelectOne(query interface{}) (*datamodels.Hero, error)
}

type HeroRepository struct {
	db *gorm.DB
}

func NewHeroRepository() IHeroRepository {
	return &HeroRepository{db: datasource.PqDB.DB}
}

func (h *HeroRepository) SelectOne(query interface{}) (*datamodels.Hero, error) {
	var hero *datamodels.Hero
	if err := h.db.Where(query).Model(hero).Error; err != nil {
		return nil, err
	}

	if hero == nil {
		return nil, errors.New("错误信息")
	}

	return hero, nil
}
