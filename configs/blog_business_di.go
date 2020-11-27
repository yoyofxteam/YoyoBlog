package configs

import (
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"yoyoFxBlog/repository"
	"yoyoFxBlog/service"
)

func ConfigBusiness(serviceCollection *DependencyInjection.ServiceCollection) {
	serviceCollection.AddTransient(repository.NewBaseRepository)
	serviceCollection.AddTransient(service.NewBlogService)
}
