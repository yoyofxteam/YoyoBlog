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
	*Mvc.RequestBody
	Id           int    `param:"id"`
	Title        string `param:"title"`        //标题
	Introduction string `param:"introduction"` //简介
	Content      string `param:"content"`      //内容
	ViewCount    int    `param:"viewCount"`    //浏览次数
}

type PageRequest struct {
	*Mvc.RequestBody
	PageIndex int `param:"pageIndex"`
	PageSize  int `param:"pageSize"`

}



func (controller *BlogController) AddBlog(blog *BlogRequest) ActionResult.IActionResult {
	data := controller.blogService.AddLog(domain.Blog{Id: blog.Id, Title: blog.Title, Introduction: blog.Introduction, Content: blog.Content, ViewCount: blog.ViewCount})
	return ActionResult.Json{Data: data}
}

func (controller *BlogController) GetBlogList(PageRequest *PageRequest) ActionResult.IActionResult {
	data := controller.blogService.QueryBlogList(PageRequest.PageIndex, PageRequest.PageSize)
	return ActionResult.Json{Data: data}
}

func (controller *BlogController) BlogList(PageRequest *PageRequest) ActionResult.IActionResult {
	data := controller.blogService.QueryBlogList(PageRequest.PageIndex, PageRequest.PageSize)
	return ActionResult.Json{Data: data}
}
