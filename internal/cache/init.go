package cache

import (
	"github.com/go-redis/redis"
	"mall_go/pkg/setting"
)

// 初始化连接

func NewRedisClient(redisSet *setting.RedisSettings) (*redis.Client, error) {
	RDB := redis.NewClient(&redis.Options{
		Addr:     redisSet.Host,
		Password: redisSet.Password, // no password set
		DB:       0,                 // use default DB
	})
	return RDB, nil
}
func SelectDB(db int) {
	//RDB.Do("SELECT", db)
}
