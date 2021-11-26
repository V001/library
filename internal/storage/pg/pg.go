package pg

import (
	"fmt"
	"github.com/v001/library/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// TODO add conn string from cfg
func Dial(cfg *configs.Config) (*gorm.DB, error) {
	time.Sleep(25 * time.Second)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB.User, cfg.DB.Password, cfg.DB.URL, cfg.DB.DBName)
	//dsn := "root:root@tcp(localhost:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
