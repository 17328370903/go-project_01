package conf

import (
	"fmt"
	"project01/global"
	"project01/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化mysql 连接
func InitDb() *gorm.DB {
	dsn := viper.GetString("mysql.dns")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		global.Logger.Error(fmt.Sprintf("mysql connect error: %s", err.Error()))
		panic(fmt.Sprintf("mysql connect error: %s", err.Error()))
	}

	db.AutoMigrate(models.User{})

	return db
}
