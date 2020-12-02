package repository_impl

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yoyofx/yoyogo/Abstractions"
	"strings"
)

type BaseRepository struct {
	config Abstractions.IConfiguration
}

func NewBaseRepository(config Abstractions.IConfiguration) *BaseRepository {
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
