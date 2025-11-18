package dao

import (
	"context"

	"github.com/extra-time-zone/xgin/database"
	"github.com/redis/go-redis/v9"
)

type rdConn struct {
	ctx context.Context
	*redis.ClusterClient
}

func NewRedis(ctx context.Context) *rdConn {
	return &rdConn{
		ctx,
		database.GetRD(),
	}
}
