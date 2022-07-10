// Package core
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 服务器核心
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2022/7/3 18:35
 */
package core

import (
	"gin-demo/global"
	"gin-demo/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 初始化路由
	Router := initialize.Routers()
	// 初始化服务
	s := initServer(global.Config.System.Part, Router)
	// 启动服务器
	s.ListenAndServe().Error()
}
