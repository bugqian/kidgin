package Models

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"kidgin/Config"
)

var db *gorm.DB

var RedisClient *redis.Client

func init() {
	RegisterGormDB()
	//ConnectRedis()
}

func RegisterGormDB() {
	mysqlCon := Config.Config("database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", mysqlCon["root"], mysqlCon["passward"], mysqlCon["host"], mysqlCon["port"], mysqlCon["db"], "Asia%2FShanghai")
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}, Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}

func ConnectRedis() {
	redisCon := Config.Config("redis")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisCon["host"], redisCon["port"]),
		Password: redisCon["auth"],
		DB:       0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
}
