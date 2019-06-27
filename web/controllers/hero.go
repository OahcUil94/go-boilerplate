package controllers

import (
	"github.com/OahcUil94/go-boilerplate/models"
	"github.com/OahcUil94/go-boilerplate/services"
	"github.com/kataras/iris"
)

type HeroController struct {
	Ctx iris.Context
	Service services.IHeroService
}

func (h *HeroController) GetBy(name string) {
	_, _ = h.Ctx.JSON(map[string]string{"a": "b"})
}

func (h *HeroController) PostCreate() {
	var data models.HeroCreateBody
	if err := h.Ctx.ReadJSON(&data); err != nil {
		_, _ = h.Ctx.JSON(&models.Result{
			Code: 400,
			Msg: "参数解析错误",
		})
		return
	}

	_, _ = h.Ctx.JSON(h.Service.CreateHero(&data))
}
