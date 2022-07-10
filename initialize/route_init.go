// Package initialize
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 初始化路由
 * @File:  route_init
 * @Version: 1.0.0
 * @Date: 2022/7/3 18:39
 */
package initialize

import (
	"gin-demo/middleware"
	"gin-demo/router/base"
	"github.com/gin-gonic/gin"
)

// Routers 配置路由，依赖gin
// 日志、跨域、JWT
// public的路由不需要鉴权 admin/login不需要鉴权（处于登录态才颁发JWT）
func Routers() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Logger())
	engine.Use(middleware.Cors())

	root := engine.Group("api")
	{
		base.InitBaseRouter(root)
	}
	return engine
}
