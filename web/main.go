package main

import (
	"github.com/OahcUil94/go-boilerplate/config"
	"github.com/OahcUil94/go-boilerplate/services"
	"github.com/OahcUil94/go-boilerplate/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.Default()

	mvc.Configure(app.Party("/heroes"), heroes)

	_ = app.Run(iris.Addr(":" + config.Server.Port))
}

func heroes(app *mvc.Application) {
	heroService := services.NewHeroService()
	app.Register(heroService)

	app.Handle(new(controllers.HeroController))
}
