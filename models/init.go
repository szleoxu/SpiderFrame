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
	cfg := config.GetConfigFile()
	username, _ := cfg.GetValue("mysql", "username")
	password, _ := cfg.GetValue("mysql", "password")
	host, _ := cfg.GetValue("mysql", "host")
	port, _ := cfg.GetValue("mysql", "port")
	database, _ := cfg.GetValue("mysql", "database")
	connectionStr := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
	sqlDB, err := sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatalf("connect mysql is fail.%s", err.Error())
		os.Exit(1)
	}
	errPing:=sqlDB.Ping()
	if errPing!=nil{
		log.Fatalf("ping mysql is fail.%s", err.Error())
		os.Exit(1)
	}
	return sqlDB
}

func InitRedis(dbIndex int) *redis.Client {
	cfg := config.GetConfigFile()
	ip,_ := cfg.GetValue("redis", "host")
	pwd,_ := cfg.GetValue("redis", "password")
	rdb := redis.NewClient(&redis.Options{
		Addr:     ip,
		Password: pwd,
		DB:       dbIndex, // use default DB
	})
	return rdb
}


