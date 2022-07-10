// Package base
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 基础接口
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2022/7/3 18:44
 */
package base

import (
	baseController "gin-demo/api/base"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	userController := baseController.NewUserController()
	{
		BaseRouter.POST("hello", userController.HelloWorld)
	}
	return BaseRouter
}
