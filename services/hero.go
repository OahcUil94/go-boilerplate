package services

import (
	"fmt"
	"github.com/OahcUil94/go-boilerplate/datamodels"
	"github.com/OahcUil94/go-boilerplate/models"
	"github.com/OahcUil94/go-boilerplate/repositories"
	"strings"
	"time"
)

type IHeroService interface {
	GetHero(string) (*datamodels.Hero, error)
	CreateHero(*models.HeroCreateBody) *models.Result
}

type HeroService struct {
	repo repositories.IHeroRepository
}

func NewHeroService() IHeroService {
	return &HeroService{repo: repositories.NewHeroRepository()}
}

func (h *HeroService) GetHero(name string) (*datamodels.Hero, error) {
	return h.repo.SelectOne("name = ?", name)
}

func (h *HeroService) CreateHero(data *models.HeroCreateBody) *models.Result {
	if data == nil {
		return &models.Result{
			Code: 500,
			Msg: "参数传递错误",
		}
	}

	data.Name = strings.TrimSpace(data.Name)
	data.RealName = strings.TrimSpace(data.RealName)

	if data.Name == "" || data.RealName == "" {
		return &models.Result{
			Code: 400,
			Msg: "英雄的名字和本人姓名错误",
		}
	}

	hero, err := h.repo.SelectOne("name = ? OR real_name = ?", data.Name, data.RealName)
	fmt.Println(err)
	if err != nil {
		return &models.Result{
			Code: 500,
			Msg: "数据库错误, 获取hero失败",
		}
	}

	if hero != nil {
		return &models.Result{
			Code: 400,
			Msg: "该英雄已存在",
		}
	}

	hero = &datamodels.Hero{
		ID: time.Now().UnixNano(),
		Name: data.Name,
		RealName: data.RealName,
	}

	if err := h.repo.Create(hero); err != nil {
		return &models.Result{
			Code: 500,
			Msg: "数据库错误, 英雄创建失败",
		}
	}

	return &models.Result{
		Code: 200,
		Msg: "英雄创建成功",
		Data: *hero,
	}
}