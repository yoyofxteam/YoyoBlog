package domain

import (
	"time"
)

type Blog struct {
	Id           int
	Title        string //标题
	Introduction string //简介
	Content      string //内容
	ViewCount    int    //浏览次数
	Author       string //作者
	CreationDate *time.Time
}
