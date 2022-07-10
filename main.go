// Package main
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/10 12:52
 */
package main

import (
	"fmt"
	"gin-demo/core"
	"gin-demo/global"
	"gin-demo/initialize"
	"gin-demo/utils"
	"os"
	"strconv"
	"strings"
)

// using env:   export GIN_MODE=release
func main() {
	if len(os.Args) > 1 {
		for idx, arg := range os.Args {
			fmt.Println("参数"+strconv.Itoa(idx)+" : ", arg)
		}
		arg := strings.Split(os.Args[1], "=")
		if len(arg) >= 2 && arg[0] == "--config" {
			// 指定yaml文件的路径
			global.Config = initialize.InitProdConfig(arg[1])
		}
	} else {
		global.Config = initialize.InitDevConfig()
	}
	utils.RecordRunPid()
	initialize.InitLogger()
	core.RunWindowsServer()
}
