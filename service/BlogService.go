package service

import (
	"fmt"
	"yoyoFxBlog/domain"
	"yoyoFxBlog/repository"
)

type BlogService struct {
	baseRepository repository.BaseRepository
}

func NewBlogService(baseRepository repository.BaseRepository) *BlogService {
	return &BlogService{baseRepository: baseRepository}
}

func (service *BlogService) AddLog(blog domain.Blog) domain.Blog {

	conn := service.baseRepository.InitDBConn()
	stmt, err := conn.Prepare("INSERT INTO `blog` SET title=?,introduction=?,content=?")
	fmt.Println(err)
	res, err := stmt.Exec(blog.Title, blog.Introduction, blog.Content)
	fmt.Println(err)
	id, err := res.LastInsertId()
	blog.Id = int(id)
	return blog
}
