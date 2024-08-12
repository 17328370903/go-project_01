package main

import (
	"project01/cmd"
)

// @title 后台管理
// @version 1.0 版本
// @description 描述
func main() {

	defer cmd.Clean()
	cmd.Start()

}
