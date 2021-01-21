package domain

import "time"

type User struct {
	Id           int
	UserName     string
	NickName     string
	Password     string
	CreationDate time.Time
}


