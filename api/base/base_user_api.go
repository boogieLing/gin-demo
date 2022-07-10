// Package base
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 用户相关操作的api
 * @File:  base_user_api
 * @Version: 1.0.0
 * @Date: 2022/7/4 15:02
 */
package base

import (
	"gin-demo/models/views"
	"gin-demo/service"
	"gin-demo/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{service.NewUserService()}
}

// HelloWorld 你好世界
func (u *UserController) HelloWorld(c *gin.Context) {
	var params views.HelloVo
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, msg.NewMsg().Failed("参数异常"))
		return
	}
	Login, err := u.userService.HelloWorld(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.NewMsg().Failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK, msg.NewMsg().Success(Login))
}
