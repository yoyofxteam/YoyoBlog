package configs

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"yoyoFxBlog/controller"
)

func ConfigController(builder *mvc.ControllerBuilder) {
	builder.AddController(controller.NewUserController)
}
