package services

import (
	"fmt"
	"github.com/OahcUil94/go-boilerplate/datamodels"
	"github.com/OahcUil94/go-boilerplate/repositories"
)

type IHeroService interface {
	GetHero(string) (*datamodels.Hero, error)
}

type HeroService struct {
	repo repositories.IHeroRepository
}

func NewHeroService() IHeroService {
	return &HeroService{repo: repositories.NewHeroRepository()}
}

func (h *HeroService) GetHero(name string) (*datamodels.Hero, error) {
	return h.repo.SelectOne(fmt.Sprintf(`name = '%s'`, name))
}
