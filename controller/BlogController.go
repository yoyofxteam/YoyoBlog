package controller

import (
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
	"yoyoFxBlog/domain"
	"yoyoFxBlog/service"
)

type BlogController struct {
	*Mvc.ApiController
	blogService *service.BlogService
}

func NewBlogController(blogService *service.BlogService) *BlogController {
	return &BlogController{blogService: blogService}
}

type BlogRequest struct {
	Mvc.RequestBody
	Id           int
	Title        string `param:title`        //标题
	Introduction string `param:introduction` //简介
	Content      string `param:content`      //内容
	ViewCount    int    //浏览次数
}

func (this *BlogController) AddBlog(blog domain.Blog) ActionResult.IActionResult {
	data := this.blogService.AddLog(blog)
	return ActionResult.Json{Data: data}
}
