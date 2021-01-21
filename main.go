package main

import (
	"github.com/yoyofx/yoyogo/abstractions"
	WebApplication "github.com/yoyofx/yoyogo/web"
	"yoyoFxBlog/configs"
)

func main() {
	webHost := CreateYoyoBlogBuilder().Build()
	webHost.Run()
}

func CreateYoyoBlogBuilder() *abstractions.HostBuilder {
	configuration := abstractions.NewConfigurationBuilder().AddYamlFile("config").Build()
	return WebApplication.NewWebHostBuilder().
		UseConfiguration(configuration).
		Configure(func(app *WebApplication.WebApplicationBuilder) {
			app.UseMvc(configs.ConfigController)
		}).
		ConfigureServices(configs.ConfigBusiness)
}
