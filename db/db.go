package db

import (
	"github.com/v001/library/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection(cfg conf.Config) (*gorm.DB, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
