package redis

import (
	"context"
	"gtools/conf"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	onceRedis   sync.Once
)

func Init(ctx context.Context, redisConfig conf.Redis) {
	onceRedis.Do(func() {
		if redisClient == nil {
			redisClient = redis.NewClient(&redis.Options{
				Addr:     redisConfig.Addr,
				Username: redisConfig.Username,
				Password: redisConfig.Password,
			})
		}

		if redisClient != nil {
			err := redisClient.Ping(ctx).Err()
			if err != nil {
				hlog.CtxErrorf(ctx, "redis 连接失败. err:%s", err)
				panic("redis 连接失败")
			}
			hlog.CtxInfof(ctx, "redis 初始化成功")
		} else {
			panic("redis 连接失败")
		}
	})

	hlog.CtxInfof(ctx, "init redis success")
}

func KeySet(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return redisClient.Set(ctx, key, value, expiration).Err()
}

func GetVal(ctx context.Context, key string) *redis.StringCmd {
	return redisClient.Get(ctx, key)
}

func KeyExists(ctx context.Context, key string) bool {
	return redisClient.Exists(ctx, key).Val() > 0
}

func DelKey(ctx context.Context, key string) error {
	err := redisClient.Del(ctx, key).Err()
	if err != nil {
		hlog.CtxErrorf(ctx, "del key fail, key:%s, err:%s", key, err.Error())
		return err
	}
	return nil
}
