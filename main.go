package main

import (
	YoyoGo "github.com/yoyofx/yoyogo/WebFramework"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
)

func main() {
	YoyoGo.CreateDefaultBuilder(func(router Router.IRouterBuilder) {
		router.GET("/info",func (ctx *Context.HttpContext) {    // 支持Group方式
			ctx.JSON(200, Context.H{"info": "ok"})
		})
	}).Build().Run()       //默认端口号 :8080
}
