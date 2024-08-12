package cmd

import (
	"fmt"
	"project01/conf"
	"project01/global"
	"project01/router"
)

func Start() {
	conf.InitConfig()
	global.Logger = conf.InitLogger()
	global.DB = conf.InitDb()
	router.InitRouter()
}

func Clean() {
	fmt.Println("clean Success")
}
