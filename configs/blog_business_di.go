package configs

import (
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"yoyoFxBlog/repository"
	"yoyoFxBlog/repository/repository_impl"
)

func ConfigBusiness(serviceCollection *DependencyInjection.ServiceCollection) {

	//瞬时
	//直接注入一个容器
	serviceCollection.AddTransient(repository_impl.NewBaseRepository)
	//为注入的容器起一个名字，类似于Spring中为 Bean命名
	serviceCollection.AddTransientByName("NewBaseRepository",repository_impl.NewBaseRepository)
	//用接口的形式进行注入，用于一个接口多种实现的情况
	serviceCollection.AddTransientByImplements(repository_impl.NewBaseRepository, new(repository.BaseRepository))

	//单例
	serviceCollection.AddSingleton(repository_impl.NewBaseRepository)
	serviceCollection.AddSingletonByName("NewBaseRepository",repository_impl.NewBaseRepository)
	serviceCollection.AddTransientByImplements(repository_impl.NewBaseRepository, new(repository.BaseRepository))

}
