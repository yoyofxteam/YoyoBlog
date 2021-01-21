package repository

import (
	"database/sql"
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"strings"
	_"github.com/go-sql-driver/mysql"
)

type BaseRepository struct {
	config abstractions.IConfiguration
	db string
}

func NewBaseRepository(config abstractions.IConfiguration) *BaseRepository {
	return &BaseRepository{config: config}
}

func (baseRepository *BaseRepository) InitDBConn() *sql.DB {
	url := fmt.Sprint(baseRepository.config.Get("yoyogo.database.url"))
	username := fmt.Sprint(baseRepository.config.Get("yoyogo.database.username"))
	password := fmt.Sprint(baseRepository.config.Get("yoyogo.database.password"))
	var sb = strings.Builder{}
	sb.WriteString(username)
	sb.WriteString(":")
	sb.WriteString(password)
	sb.WriteString("@")
	sb.WriteString(url)
	connStr := sb.String()
	fmt.Println(connStr)
	conn, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
	}
	return conn
}
