package main

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/DependencyInjection"
	YoyoGo "github.com/yoyofx/yoyogo/WebFramework"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
	"yoyoFxBlog/controller"
	"yoyoFxBlog/repository/repository_impl"
	"yoyoFxBlog/service"
)

func main() {
	webHost := CreateYoyoBlogBuilder().Build()
	webHost.Run()
}

func CreateYoyoBlogBuilder() *Abstractions.HostBuilder {
	configuration := Abstractions.NewConfigurationBuilder().AddYamlFile("config").Build()
	return YoyoGo.NewWebHostBuilder().
		UseConfiguration(configuration).
		Configure(func(app *YoyoGo.WebApplicationBuilder) {
			app.UseMvc(func(builder *Mvc.ControllerBuilder) {
				builder.AddController(controller.NewBlogController)
			})
		}).
		ConfigureServices(func(serviceCollection *DependencyInjection.ServiceCollection) {
			serviceCollection.AddTransient(repository_impl.NewBaseRepository)
			serviceCollection.AddTransient(service.NewBlogService)
		})
}
