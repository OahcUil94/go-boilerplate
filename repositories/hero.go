package repositories

import (
	"errors"
	"github.com/OahcUil94/go-boilerplate/datamodels"
	"github.com/OahcUil94/go-boilerplate/datasource"
	"github.com/jinzhu/gorm"
)

type IHeroRepository interface {
	SelectOne(interface{}, ...interface{}) (*datamodels.Hero, error)
	Create(*datamodels.Hero) error
}

type HeroRepository struct {
	db *gorm.DB
}

func NewHeroRepository() IHeroRepository {
	return &HeroRepository{db: datasource.PqDB.DB}
}

func (h *HeroRepository) SelectOne(query interface{}, args ...interface{}) (*datamodels.Hero, error) {
	var hero *datamodels.Hero
	if err := h.db.Where(query, args...).Model(hero).Error; err != nil {
		return nil, err
	}

	return hero, nil
}

func (h *HeroRepository) Create(hero *datamodels.Hero) error {
	if hero == nil {
		return errors.New("hero is nil")
	}

	return h.db.Create(hero).Error
}