package models

import "gorm.io/gorm"

type User struct {
	Username string `gorm:"default:'';comment:'用户名称'"`
	Email    string `gorm:"default:'';comment:'邮箱'"`
	Password string `gorm:"default:'';comment:'密码'"`
	Pid      int64  `gorm:"comment:'父级id0顶级其他子级'"`
	gorm.Model
}
