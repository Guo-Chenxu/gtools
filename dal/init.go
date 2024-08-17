package dal

import (
	"context"
	"gtools/conf"
	"gtools/dal/mysql"
	"gtools/dal/redis"
)

func Init() {
	ctx := context.Background()
	redis.Init(ctx, conf.GetRedis())
	mysql.Init(ctx, conf.GetMysql())
}
