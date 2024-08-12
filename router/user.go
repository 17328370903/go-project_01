package router

import (
	"github.com/gin-gonic/gin"
	"project01/api"
)

func InitUserRouter() func(publicGroup *gin.RouterGroup, authGroup *gin.RouterGroup) {
	return func(publicGroup *gin.RouterGroup, authGroup *gin.RouterGroup) {
		userApi := api.NewUserApi()

		userGroup := authGroup.Group("/user")
		{
			userGroup.POST("/", userApi.Add)
		}
	}
}
