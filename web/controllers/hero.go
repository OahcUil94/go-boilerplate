package controllers

import (
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
