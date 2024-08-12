package api

import (
	"github.com/gin-gonic/gin"
	"project01/service/dto"
	"project01/utils"
)

type BaseApi struct {
}

type BindRequestDataOption struct {
	Ctx        *gin.Context
	Data       any
	IsUrlParam bool
}

type ValidateOption struct {
	Ctx *gin.Context
	Dto dto.Dto
}

// BindRequestData 自动绑定数据到结构体上
func (b BaseApi) BindRequestData(option BindRequestDataOption) error {
	var err error
	if option.IsUrlParam {
		err = option.Ctx.ShouldBindUri(&option.Data)
	} else {
		err = option.Ctx.ShouldBind(&option.Data)
	}
	if err != nil {
		b.Fail(option.Ctx, utils.ResponseJson{
			Code: 400,
			Msg:  err.Error(),
		})
	}
	return err
}

// Success 成功返回
func (b BaseApi) Success(c *gin.Context, res utils.ResponseJson) {
	utils.Success(c, res)
}

// Fail 错误返回
func (b BaseApi) Fail(c *gin.Context, res utils.ResponseJson) {
	utils.Fail(c, res)
}

func (b BaseApi) Validator(option ValidateOption) error {
	var err error
	if err := option.Dto.Validator(); err != nil {
		b.Fail(option.Ctx, utils.ResponseJson{
			Code: 422,
			Msg:  err.Error(),
		})
		return err
	}
	return err
}
