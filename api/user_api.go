package api

import (
	"github.com/gin-gonic/gin"
	"project01/service"
	"project01/service/dto"
	"project01/utils"
)

// ErrorAddUser 响应错误码
var (
	ErrorAddUser int = 10001
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		Service: service.NewUserService(),
	}
}

// Add 添加用户
func (userApi UserApi) Add(c *gin.Context) {

	var err error
	var AddUserDto dto.AddUserDTO
	err = userApi.BindRequestData(BindRequestDataOption{
		Ctx:  c,
		Data: &AddUserDto,
	})
	if err != nil {
		return
	}

	if err = userApi.Validator(ValidateOption{Ctx: c, Dto: AddUserDto}); err != nil {
		return
	}

	_, err = userApi.Service.AddUser(AddUserDto)
	if err != nil {
		userApi.Fail(c, utils.ResponseJson{Code: ErrorAddUser, Msg: err.Error()})
		return
	}

	userApi.Success(c, utils.ResponseJson{Code: 200, Msg: "添加成功"})
}
