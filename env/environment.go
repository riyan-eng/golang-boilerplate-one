package env

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

func LoadEnvironmentFile() {
	viper.SetConfigFile("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("environment: %v", err.Error())
		os.Exit(1)
	}
	fmt.Println("environment: load successfully")
}

var (
	ENV                       string
	SERVER_PORT               string
	SERVER_HOST               string
	SERVER_HOST_BE            string
	SERVER_HOST_FE            string
	SERVER_TIMEZONE           string
	JWT_SECRET_ACCESS         string
	JWT_SECRET_REFRESH        string
	JWT_SECRET_RESET          string
	JWT_EXPIRED_LOGOFF        time.Duration
	JWT_EXPIRED_ACCESS        time.Duration
	JWT_EXPIRED_REFRESH       time.Duration
	JWT_EXPIRED_RESET         time.Duration
	SMTP_HOST                 string
	SMTP_PORT                 string
	SMTP_EMAIL                string
	SMTP_PASSWORD             string
	POSTGRE_HOST              string
	POSTGRE_PORT              string
	POSTGRE_TIMEZONE          string
	POSTGRE_USERNAME          string
	POSTGRE_PASSWORD          string
	POSTGRE_DATABASE          string
	POSTGRE_CONN_MAX_IDLE     int
	POSTGRE_CONN_MAX_OPEN     int
	POSTGRE_CONN_MAX_LIFETIME time.Duration
	REDIS_HOST                string
	REDIS_PORT                string
	REDIS_USERNAME            string
	REDIS_PASSWORD            string
	REDIS_DATABASE            int
)

func NewEnvironment() {
	ENV = viper.GetString("env")
	SERVER_HOST = viper.GetString("server.host")
	SERVER_PORT = viper.GetString("server.port")
	SERVER_HOST_BE = viper.GetString("server.host_be")
	SERVER_HOST_FE = viper.GetString("server.host_fe")
	SERVER_TIMEZONE = viper.GetString("server.timezone")
	JWT_SECRET_ACCESS = viper.GetString("jwt.secret.access")
	JWT_SECRET_REFRESH = viper.GetString("jwt.secret.refresh")
	JWT_SECRET_RESET = viper.GetString("jwt.secret.reset")
	JWT_EXPIRED_LOGOFF = viper.GetDuration("jwt.expired.logoff")
	JWT_EXPIRED_ACCESS = viper.GetDuration("jwt.expired.access")
	JWT_EXPIRED_REFRESH = viper.GetDuration("jwt.expired.refresh")
	JWT_EXPIRED_RESET = viper.GetDuration("jwt.expired.reset")
	SMTP_HOST = viper.GetString("smtp.host")
	SMTP_PORT = viper.GetString("smtp.port")
	SMTP_EMAIL = viper.GetString("smtp.email")
	SMTP_PASSWORD = viper.GetString("smtp.password")
	POSTGRE_HOST = viper.GetString("postgre.host")
	POSTGRE_PORT = viper.GetString("postgre.port")
	POSTGRE_TIMEZONE = viper.GetString("postgre.timezone")
	POSTGRE_USERNAME = viper.GetString("postgre.username")
	POSTGRE_PASSWORD = viper.GetString("postgre.password")
	POSTGRE_DATABASE = viper.GetString("postgre.database")
	POSTGRE_CONN_MAX_IDLE = viper.GetInt("postgre.connection.max_idle")
	POSTGRE_CONN_MAX_OPEN = viper.GetInt("postgre.connection.max_open")
	POSTGRE_CONN_MAX_LIFETIME = viper.GetDuration("postgre.connection.max_lifetime")
	REDIS_HOST = viper.GetString("redis.host")
	REDIS_PORT = viper.GetString("redis.port")
	REDIS_USERNAME = viper.GetString("redis.username")
	REDIS_PASSWORD = viper.GetString("redis.password")
	REDIS_DATABASE = viper.GetInt("redis.database")
}
