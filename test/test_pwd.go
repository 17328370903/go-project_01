package main

import (
	"fmt"
	"project01/utils"
)

func main() {
	hash,_ := utils.PasswordHash("123456")
	fmt.Println(hash)
}