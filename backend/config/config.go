package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port int
	Mode string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	Charset  string
}

type JWTConfig struct {
	Secret string
	Expire int
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	AppConfig = &Config{
		Server: ServerConfig{
			Port: viper.GetInt("server.port"),
			Mode: viper.GetString("server.mode"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.dbname"),
			Charset:  viper.GetString("database.charset"),
		},
		JWT: JWTConfig{
			Secret: viper.GetString("jwt.secret"),
			Expire: viper.GetInt("jwt.expire"),
		},
	}

	log.Println("配置文件加载成功")
}
