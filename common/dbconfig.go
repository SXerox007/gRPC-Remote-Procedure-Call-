package common

import "time"

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Debug        bool
	Address      string
	Port         int
	MaxIdleConns int
	MaxOpenConns int
	User         string
	Passwd       string
	DbName       string
	Prefix       string
	PingInterval time.Duration
}

// Init 初始化数据库配置
func (db *MySQLConfig) Init() {
	if db.Address == "" {
		db.Address = "localhost"
	}
	if db.Port == 0 {
		db.Port = 5051
	}
	if db.MaxIdleConns == 0 {
		db.MaxIdleConns = 64
	}
	if db.MaxOpenConns == 0 {
		db.MaxOpenConns = 16
	}
	if db.User == "" {
		db.User = "root"
	}
	if db.PingInterval.Seconds() == 0 {
		duration, _ := time.ParseDuration("10s")
		db.PingInterval = duration
	}
}
