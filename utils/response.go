package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func bindingDefaultStatus(res ResponseJson) int {
	if res.Status == 0 {
		return http.StatusOK
	}
	return res.Status
}

func Success(c *gin.Context, res ResponseJson) {
	c.AbortWithStatusJSON(bindingDefaultStatus(res), res)
}

func Fail(c *gin.Context, res ResponseJson) {
	c.AbortWithStatusJSON(bindingDefaultStatus(res), res)
}
