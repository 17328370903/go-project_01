package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"project01/docs"
	"project01/global"
	"project01/middleware"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var initRoutesFn []func(*gin.RouterGroup, *gin.RouterGroup)

func InitRouter() {

	//优雅处理退出
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	// docs.SwaggerInfo.BasePath = "/api/v1"

	r := gin.Default()
	//记录recover 错误 panic

	docs.SwaggerInfo.BasePath = "/api/v1"
	//swagger 日志路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(middleware.Cors(), middleware.Recover)

	authGroup := r.Group("/api/v1")
	publicGroup := r.Group("/api/v1/public")
	loadRouter()

	for _, fn := range initRoutesFn {
		fn(publicGroup, authGroup)
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Error(fmt.Sprintf("服务启动失败,原因:%s", err.Error()))
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("服务关闭失败,原因:%s", err.Error()))
		return
	}

	fmt.Println("server 关闭成功")

}

func loadRouter() {
	initRoutesFn = append(initRoutesFn, InitUserRouter())
}
