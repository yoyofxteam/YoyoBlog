package service

import (
	"fmt"
	"yoyoFxBlog/domain"
	"yoyoFxBlog/repository/repository_impl"
)

type BlogService struct {
	baseRepository *repository_impl.BaseRepository
}

func NewBlogService(baseRepository *repository_impl.BaseRepository) *BlogService {
	return &BlogService{baseRepository: baseRepository}
}

func (service *BlogService) AddLog(blog domain.Blog) domain.Blog {

	conn := service.baseRepository.InitDBConn()
	defer conn.Close()
	stmt, err := conn.Prepare("INSERT INTO `blog` SET title=?,introduction=?,content=?")
	fmt.Println(err)
	res, err := stmt.Exec(blog.Title, blog.Introduction, blog.Content)
	fmt.Println(err)
	id, err := res.LastInsertId()
	blog.Id = int(id)
	return blog
}

func (service *BlogService) QueryBlogList(pageIndex int, pageSize int) domain.Page {
	conn := service.baseRepository.InitDBConn()
	defer conn.Close()
	res := domain.Page{}
	rows, err := conn.Query("SELECT COUNT(0) as count FROM `blog` ")
	if err != nil {

	}
	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		res.TotalCount = count
	}
	start := (pageIndex - 1) * pageSize
	sql := fmt.Sprintf("SELECT *FROM `blog` ORDER BY creation_date LIMIT %d,%d", start, pageSize)
	rows, err = conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	blogList := make([]domain.Blog, 0)
	for rows.Next() {
		element := domain.Blog{}
		err := rows.Scan(&element.Id, &element.Title, &element.Introduction, &element.Content, &element.ViewCount, &element.Author, &element.CreationDate)
		if err != nil {
			continue
		}
		blogList = append(blogList, element)

	}
	res.Data = blogList
	return res
}
