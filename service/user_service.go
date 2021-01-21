package service

import (
	"crypto/md5"
	"fmt"
	"yoyoFxBlog/domain"
	"yoyoFxBlog/repository"
	"yoyoFxBlog/utils"
)

type UserService struct {
	repository *repository.BaseRepository
}

func  NewUserService(baseRepository *repository.BaseRepository)*UserService  {
	return &UserService{repository:baseRepository}
}

func (userService *UserService) AddUser(user domain.User) domain.ResultModel {
	var res = domain.ResultModel{}
	if user.Password == "" {
		panic("密码不能为空")
		res.Success = false
		res.Message = "密码不能为空"
		return res
	}
	cipherStr:= md5.Sum([]byte(user.Password))
	md5Pwd := fmt.Sprintf("%x",cipherStr)
	conn := userService.repository.InitDBConn()
	defer conn.Close()
	stmt, err := conn.Prepare("INSERT INTO t_user VALUE (0,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	insertRes, err := stmt.Exec(user.UserName, user.NickName, md5Pwd, user.CreationDate)
	if err != nil {
		panic(err)
	}
	actRows, _ := insertRes.RowsAffected()
	if actRows > 0 {
		res.Success = true
		res.Message = "添加成功"
		res.Data = user
		return res
	}
	res.Success = false
	res.Message = "添加用户失败"
	return res
}

func (userService *UserService) Login(username string, password string) domain.ResultModel {
	var res = domain.ResultModel{}
	conn := userService.repository.InitDBConn()
	defer conn.Close()
	stmt, err := conn.Prepare("select * from t_user where username=?")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(username)
	if err != nil {
		panic(err)
	}
	var userInfo = domain.User{}
	for rows.Next() {
		err := rows.Scan(&userInfo.Id, &userInfo.UserName, &userInfo.NickName, &userInfo.Password, &userInfo.CreationDate)
		if err != nil {
			panic(err)
		}
	}

	if !utils.IsNullOrEmpty(userInfo.UserName) {
		panic("没有找到对应的数据")
	}

	md5Pwd := fmt.Sprintf("%x",md5.Sum([]byte(password)))
	if md5Pwd != userInfo.Password {
		res.Message = "密码不正确"
		res.Success = false
		return res
	}
	userInfo.Password = ""
	res.Message = "登录成功"
	res.Success = true
	res.Data = userInfo
	return res
}
