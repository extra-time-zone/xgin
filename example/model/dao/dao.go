package dao

import (
	"github.com/extra-time-zone/xgin/database"
	"github.com/extra-time-zone/xgin/example/config"
)

func Init() {
	// mysql
	if err := initMysql(); err != nil {
		panic(err)
	}

	// redis
	if err := initRedis(); err != nil {
		panic(err)
	}
}

func initMysql() error {
	dbConf := make([]*database.MySqlConfig, 0, len(config.Config.MySql))
	for _, v := range config.Config.MySql {
		//master db
		mdb := makeDBConfig(v)
		//slave db
		sdb := make([]*database.MySqlConfig, 0, len(v.Slave))
		for _, s := range v.Slave {
			sdb = append(sdb, makeDBConfig(s))
		}

		mdb.Slave = sdb
		dbConf = append(dbConf, mdb)
	}

	if len(dbConf) > 0 {
		return database.InitDB(dbConf)
	}
	return nil
}

func initRedis() error {
	rdConf := makeRDConfig(config.Config.Redis)
	if config.Config.Redis.Tls {
		rdConf.TLS = true
		rdConf.ServerName = config.Config.Redis.ServerName
	}
	return database.InitRedis(rdConf)
}

func Close() {
	database.CloseDB()
	database.CloseRD()
}

func makeDBConfig(conf *config.MySqlConfig) *database.MySqlConfig {
	return &database.MySqlConfig{
		Role:         conf.Role,
		Host:         conf.Host,
		Port:         conf.Port,
		Dbname:       conf.Dbname,
		Username:     conf.Username,
		Password:     conf.Password,
		Charset:      conf.Charset,
		Collation:    conf.Collation,
		MaxIdleConns: conf.MaxIdleConns,
		MaxOpenConns: conf.MaxOpenConns,
		MaxLifetime:  conf.MaxLifetime,
		MaxIdleTime:  conf.MaxIdleTime,
		Slave:        nil,
	}
}

func makeRDConfig(conf *config.RedisConfig) *database.RedisClusterConfig {
	return &database.RedisClusterConfig{
		Host:         conf.Host,
		Username:     conf.Username,
		Password:     conf.Password,
		MaxIdleConns: conf.MaxIdleConns,
		MinIdleConns: conf.MinIdleConns,
		PoolSize:     conf.PoolSize,
		TLS:          false,
		ServerName:   "",
	}
}
