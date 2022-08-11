package easy_serve

import (
	"context"
	"fmt"

	"github.com/Sunqi43797189/easy_serve/config"
	"github.com/go-redis/redis/v8"
)

var redisMap = make(map[string]*redisObj)

type redisObj struct {
	client *redis.Client
	err   error
}

func initRedis() {
	for _, conf := range config.C.Redis {
		rdb := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
			Username: conf.UserName,
			Password: conf.Password, // no password set
			DB:       conf.DB,       // use default DB
		})
		var obj redisObj
		if rdb == nil {
			obj.client = nil
			obj.err = fmt.Errorf("name %s not exists ", conf.Name)
		} else if err := rdb.Ping(context.Background()).Err(); err != nil {
			obj.err = err
		} else {
			obj.client = rdb
			obj.err = nil
		}
		redisMap[conf.Name] = &obj
	}
}
