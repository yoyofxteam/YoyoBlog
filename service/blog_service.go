package service

import (
	"fmt"
	"yoyoFxBlog/domain"
	"yoyoFxBlog/repository"
)

type BlogService struct {
	baseRepository *repository.BaseRepository
}

func NewBlogService(baseRepository *repository.BaseRepository) *BlogService {
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
	start := (pageIndex-1) * pageSize
	sql := fmt.Sprintf("SELECT * FROM `blog` ORDER BY creation_date LIMIT %d,%d", start, pageSize)
	rows, err = conn.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	blogList := make([]domain.Blog, 20)
	for rows.Next() {

		var id int
		var title string
		var introduction string
		var content string
		var viewCount int
		var author string

		rows.Scan(&id, &title, &introduction, &content, &viewCount, &author)
		element := domain.Blog{
			Id:           id,
			Title:        title,
			Introduction: introduction,
			Content:      content,
			ViewCount:    viewCount,
			Author:       author,
		}
		blogList = append(blogList, element)
	}
	res.Data = blogList
	return res
}
