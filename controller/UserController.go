package controller

import (
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
)

type UserInfo struct {
	Mvc.RequestBody
	UserName string `param:"username"`
	Password string `param:"password"`
}

func NewUserController()  *UserController{
	return &UserController{}
}

type UserController struct {
	*Mvc.ApiController
}

func (*UserController)Register(ctx *Context.HttpContext, request *UserInfo) ActionResult.IActionResult {
	result := Mvc.ApiResult{Success: true, Message: "ok", Data: request}
	return ActionResult.Json{Data: result}
}

