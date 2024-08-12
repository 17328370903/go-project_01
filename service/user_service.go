package service

import (
	"project01/dao"
	"project01/models"
	"project01/service/dto"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (u UserService) AddUser(addUserDto dto.AddUserDTO) (models.User, error) {
	var user models.User
	addUserDto.ToUserModel(&user)
	user.Pid = 0
	err := u.Dao.AddUser(&user)
	return user, err
}
