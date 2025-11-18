package dao

import (
	"context"
	"fmt"
	"github.com/extra-time-zone/xgin/example/config"
	"testing"
)

func init() {
	config.Config.MySql = []*config.MySqlConfig{
		{
			Host:      "54.69.237.139",
			Port:      8096,
			Dbname:    "test-sports-db",
			Username:  "root",
			Password:  "123456",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
	config.Config.Redis = &config.RedisConfig{
		Host: []string{"54.69.237.139:8084", "54.69.237.139:8085", "54.69.237.139:8086"},
	}

	if err := initMysql(); err != nil {
		panic(err)
	}
	if err := initRedis(); err != nil {
		panic(err)
	}
}

func TestRedisDev(t *testing.T) {
	//ctx := context.Background()
	//
	////dev
	//env := global.ServerRunEnvDev
	//global.ServerRunEnv = &env
	//redisConfigDev := config.RedisConfig{
	//	Host:         []string{"54.69.237.139:8084", "54.69.237.139:8085", "54.69.237.139:8086"},
	//	Username:     "",
	//	Password:     "",
	//	ServerName:   "",
	//	MaxIdleConns: 50,
	//	MinIdleConns: 50,
	//	PoolSize:     20,
	//	Db:           0,
	//}
	//
	//rdPoolConn := NewRedis(ctx)
	//fmt.Printf("rdPoolConn : %+v \n", rdPoolConn)
	//
	//name, err := rdPoolConn.rd.Get(ctx, "name").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(name)
	//
	//rdPoolConn.CreateRedisCluster(&redisConfigDev)
	//
	//rdPoolConn1 := NewRedis(ctx, &redisConfigDev)
	//fmt.Printf("rdPoolConn1: %+v \n", rdPoolConn1)
	//
	//name, err = rdPoolConn1.rd.Get(ctx, "name").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(name)
}

func TestRedisCluster(t *testing.T) {
	ctx := context.Background()

	rd := NewRedis(ctx)
	fmt.Printf("rd : %+v \n", rd)

	key := "animal"
	cmd := rd.GetRD().Get(ctx, key)
	fmt.Printf("cmd.err: %+v \n", cmd.Err())

	result, err := cmd.Result()
	fmt.Printf("cmd.result.err: %+v \n", err)
	fmt.Printf("cmd.result.val: %+v \n", result)
}
