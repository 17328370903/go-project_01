package api

import (
	"project01/models"
	"project01/utils"

	"github.com/gin-gonic/gin"
)

type ProductApi struct {
	BaseApi
}

// List
// @Tags 产品管理
// @Summary 产品列表
// @Description 产品列表
// @Param page query int true "当前页码" default(1)
// @Param page_size query int true "分页数量" default(10)
// @Router /user [get]
func (p ProductApi) List(c *gin.Context) {
	var data models.Product

	if err := p.BindRequestData(BindRequestDataOption{Ctx: c, Data: &data}); err != nil {
		return
	}
	p.Success(c, utils.ResponseJson{Data: data})
}

// Create
// @Tags 产品管理
// @Summary 添加产品
// @Description 添加产品
// @Accept json
// @Param name body string true "产品名称"
// @Param price body number true "产品价格"
// @Router /user [post]
// @Success 200 {object} response.ResponseJson "获取成功"
// @Header 200 {string} token "token"
func (p ProductApi) Create(c *gin.Context) {

}
