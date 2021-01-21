package service

import (
	"encoding/json"
	redis "github.com/yoyofx/yoyogo/pkg/cache/redis"
	redisdb "github.com/yoyofx/yoyogo/pkg/datasources/redis"
	"yoyoFxBlog/domain"
	"yoyoFxBlog/repository"
)

const RedisPrefix = "yoyofx:tod"

type HubService struct {
	repository  repository.BaseRepository
	redisSource redisdb.RedisDataSource
}

func NewHubService(repository repository.BaseRepository, redisSource redisdb.RedisDataSource) *HubService {
	return &HubService{
		repository:  repository,
		redisSource: redisSource,
	}
}

func (hub *HubService) SyncTodoList(body string)  {

}

func (hub *HubService) GetTodoList() []domain.TaskItem {
	conn := hub.repository.InitDBConn()
	defer conn.Close()
	stmt, err := conn.Prepare("select * from t_task_item")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	res := make([]domain.TaskItem, 0)
	for rows.Next() {
		var ele = domain.TaskItem{}
		err := rows.Scan(ele.Id, ele.Editable, ele.Content, ele.Point, ele.Title)
		if err != nil {
			panic(err)
		}
		res = append(res, ele)
	}
	redisConn, _, _ := hub.redisSource.Open()
	client := redisConn.(redis.IClient)
	jsonByte, err := json.Marshal(res)
	err = client.GetKVOps().Set(RedisPrefix, string(jsonByte), 0)
	if err != nil {
		panic(err)
	}
	return res
}
