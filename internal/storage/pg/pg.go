package pg

import (
	"github.com/v001/library/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO add conn string from cfg
func Dial(cfg *configs.Config) (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
