package models

import (
	"SpiderFrame/config"
	"database/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func InitMysql() *sql.DB {
	username := config.GetValue("mysql", "username")
	password := config.GetValue("mysql", "password")
	host := config.GetValue("mysql", "host")
	port := config.GetValue("mysql", "port")
	database := config.GetValue("mysql", "database")
	connectionStr := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
	sqlDB, err := sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatalf("connect mysql is fail.%s", err.Error())
		os.Exit(1)
	}
	errPing := sqlDB.Ping()
	if errPing != nil {
		log.Fatalf("ping mysql is fail.%s", err.Error())
		os.Exit(1)
	}
	return sqlDB
}

func InitRedis(dbIndex int) *redis.Client {
	ip := config.GetValue("redis", "host")
	pwd := config.GetValue("redis", "password")
	rdb := redis.NewClient(&redis.Options{
		Addr:     ip,
		Password: pwd,
		DB:       dbIndex, // use default DB
	})
	return rdb
}
