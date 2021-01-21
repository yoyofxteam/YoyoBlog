package controller

import (
	"github.com/yoyofx/yoyogo/web/mvc"
	"time"
	"yoyoFxBlog/domain"
	"yoyoFxBlog/service"
)

type UserController struct {
	*mvc.ApiController
	userService *service.UserService
}

type UserDTO struct {
	*mvc.RequestBody
	Id           int
	UserName     string `param:"userName"`
	NickName     string `param:"nickName"`
	Password     string `param:"password"`
	CreationDate  time.Time `param:creationDate`
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (controller *UserController) AddUser(user *UserDTO) domain.ResultModel {

	return controller.userService.AddUser(domain.User{UserName: user.UserName, NickName: user.NickName, Password: user.Password, CreationDate: time.Now()})
}
