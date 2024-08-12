package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username      string `gorm:"comment:用户名;default=''"`
	Email         string `gorm:"comment:邮箱;default=''"`
	Password      string `gorm:"comment:密码;default=''"`
	LastLoginTime int64  `gorm:"comment:最后登录时间;default=0"`
	LastLoginIp   string `gorm:"comment:最后登录ip;default=''"`
}
