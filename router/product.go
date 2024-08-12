package router

import (
	"project01/api"

	"github.com/gin-gonic/gin"
)

// 初始化产品路由
func InitProductRouter(publicGroup *gin.RouterGroup, AuthGroup *gin.RouterGroup) {
	productGroup := AuthGroup.Group("/product")
	productGroup.POST("/", api.ProductApi{}.List)
	productGroup.POST("/add", api.ProductApi{}.Create)
}
