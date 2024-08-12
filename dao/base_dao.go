package dao

import (
	"gorm.io/gorm"
	"project01/global"
)

type BaseDao struct {
	Orm *gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{
		Orm: global.DB,
	}
}
