package configs

import (
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"yoyoFxBlog/repository"
	"yoyoFxBlog/service"
)

func ConfigBusiness(serviceCollection *dependencyinjection.ServiceCollection) {

	//瞬时
	//直接注入一个容器
	serviceCollection.AddTransient(repository.NewBaseRepository)
	serviceCollection.AddTransient(service.NewUserService)
}
