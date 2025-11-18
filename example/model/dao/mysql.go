package dao

import (
	"context"

	"github.com/extra-time-zone/xgin/database"
	"gorm.io/gorm"
)

type dbConn struct {
	ctx context.Context
	*gorm.DB
}

func NewDBConn(ctx context.Context, roles ...string) *dbConn {
	role := database.RoleDefault
	if len(roles) > 0 {
		role = roles[0]
	}

	return &dbConn{
		ctx,
		database.GetDB(role),
	}
}
