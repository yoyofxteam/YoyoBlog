package configs

import (
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
	"yoyoFxBlog/controller"
)

func ConfigController(builder *Mvc.ControllerBuilder) {
	builder.AddController(controller.NewUserController)
	builder.AddController(controller.NewBlogController)
}
