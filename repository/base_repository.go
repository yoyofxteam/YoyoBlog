package repository

import "database/sql"

type BaseRepository interface {
	InitDBConn() *sql.DB

}