package main

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	YoyoGo "github.com/yoyofx/yoyogo/WebFramework"
	"yoyoFxBlog/configs"
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
			app.UseMvc(configs.ConfigController)
		}).
		ConfigureServices(configs.ConfigBusiness)
}
