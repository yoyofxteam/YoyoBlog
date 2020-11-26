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
	//configuration := Abstractions.NewConfigurationBuilder().AddYamlFile("config").Build()
	return YoyoGo.NewWebHostBuilder().
		//UseConfiguration(configuration).
		Configure(func(app *YoyoGo.WebApplicationBuilder) {
			app.UseMvc(func(builder *Mvc.ControllerBuilder) {
				//builder.AddViewsByConfig()
				builder.AddController(controller.NewUserController)
			})
		})
}
