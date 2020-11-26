package main

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	YoyoGo "github.com/yoyofx/yoyogo/WebFramework"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
	"yoyoFxBlog/controller"
)

func main() {

	webHost := CreateYoyoBlogBuilder().Build()
	webHost.Run()
}

func CreateYoyoBlogBuilder() *Abstractions.HostBuilder {
	return YoyoGo.NewWebHostBuilder().
		Configure(func(app *YoyoGo.WebApplicationBuilder) {
			app.UseMvc(func(builder *Mvc.ControllerBuilder) {
				builder.AddController(controller.NewUserController)
			})
		})
}
