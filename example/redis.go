package example

import (
	"SpiderFrame/models"
	"context"
	"time"
)

func Set(key string ,value string,expiration time.Duration ,dbIndex int) {
	var ctx = context.Background()
	rdb := models.InitRedis(dbIndex)
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
}
