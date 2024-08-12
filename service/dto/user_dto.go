package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"project01/models"
)

type AddUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (d AddUserDTO) Validator() error {
	err := validation.ValidateStruct(&d,
		validation.Field(&d.Username, validation.Required.Error("用户名不能为空")),
		validation.Field(&d.Email, validation.Required.Error("邮箱不能为空")),
		validation.Field(&d.Password, validation.Required.Error("密码不能为空")),
	)
	return GetValidateError(err)
}

func (d AddUserDTO) ToUserModel(user *models.User) {
	user.Username = d.Username
	user.Password = d.Password
	user.Email = d.Email
}
