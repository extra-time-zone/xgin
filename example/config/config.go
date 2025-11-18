package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigStruck struct {
	Http  *HttpConfig
	MySql []*MySqlConfig
	Redis *RedisConfig
}

type HttpConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Env  string `json:"env"`
	Mode string `json:"mode"`
}

type MySqlConfig struct {
	Role         string         `json:"role"`
	Host         string         `json:"host"`
	Port         int            `json:"port"`
	Dbname       string         `json:"dbname"`
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	Charset      string         `json:"charset"`
	Collation    string         `json:"collation"`
	MaxIdleConns int            `json:"max_idle_conns"`
	MaxOpenConns int            `json:"max_open_conns"`
	MaxLifetime  int            `json:"max_lifetime"`
	MaxIdleTime  int            `json:"max_idle_time"`
	Slave        []*MySqlConfig `json:"slave"`
}

type RedisConfig struct {
	Host         []string `json:"host"`
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	ServerName   string   `json:"server_name"`
	MaxIdleConns int      `json:"max_idle_conns"` //最大空闲链接数
	MinIdleConns int      `json:"min_idle_conns"` //最小空闲链接数
	PoolSize     int      `json:"pool_size"`      //链接池最大链接数
	Prefix       string   `json:"prefix"`         //key前缀
	Tls          bool     `json:"tls"`
}

// Config data
var Config = ConfigStruck{}

// Init 初始化配置信息
func Init() {
	// file xglobal
	if err := parseConfigFromToml(&Config); err != nil {
		panic(fmt.Sprintf("xglobal init from file error:%v", err))
	}

	fmt.Printf("config.Http: %+v\n", Config.Http)
}

// 获取配置文件并解析到指定的变量
func parseConfigFromToml(config any) error {
	viper.SetConfigFile("./conf/config_local.toml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	//parse
	return viper.Unmarshal(config)
}
