package configs

import (
	"fmt"
	"sync"
)

const (
	TextFormat = LogFormat("text")
	JsonFormat = LogFormat("json")
)

var config *Config

type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}
type App struct {
	Name      string `toml:"name" env:"APP_NAME"`
	Host      string `toml:"host" env:"APP_HOST"`
	Port      string `toml:"port" env:"APP_PORT"`
	Key       string `toml:"key" env:"APP_KEY"`
	EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

type LogTo string
type LogFormat string

// Log todo
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

// MySQL todo
type MySQL struct {
	Host        string `toml:"host" env:"MYSQL_HOST"`
	Port        string `toml:"port" env:"MYSQL_PORT"`
	UserName    string `toml:"username" env:"MYSQL_USERNAME"`
	Password    string `toml:"password" env:"MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
	lock        sync.Mutex
}

func newDefaultAPP() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
		Key:  "default",
	}
}
func newDefaultLog() *Log {
	return &Log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
}
func newDefaultMYSQL() *MySQL {
	return &MySQL{
		Database:    "go_course",
		Host:        "127.0.0.1",
		Port:        "3306",
		MaxOpenConn: 200,
		MaxIdleConn: 50,
		MaxLifeTime: 1800,
		MaxIdleTime: 600,
	}
}
func newConfig() *Config {
	return &Config{
		App:   newDefaultAPP(),
		Log:   newDefaultLog(),
		MySQL: newDefaultMYSQL(),
	}
}
func NewDefaultConfig() *Config {
	return &Config{
		App:   newDefaultAPP(),
		Log:   newDefaultLog(),
		MySQL: newDefaultMYSQL(),
	}
}
func (a *App) HttpAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

// 连接池, driverConn具体的连接对象, 他维护着一个Socket
// pool []*driverConn, 维护pool里面的连接都是可用的, 定期检查我们的conn健康情况
// 某一个driverConn已经失效, driverConn.Reset(), 清空该结构体的数据, Reconn获取一个连接, 让该conn借壳存活
// 避免driverConn结构体的内存申请和释放的一个成本
