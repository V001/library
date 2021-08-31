package db

import (
	"github.com/v001/library/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB(conf *configs.DBConf) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(conf.URL), &gorm.Config{})
}
