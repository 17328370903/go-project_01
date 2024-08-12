package dao

import (
	"project01/models"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{NewBaseDao()}
	}
	return userDao
}

func (u UserDao) AddUser(user *models.User) error {
	return u.Orm.Create(&user).Error
}
