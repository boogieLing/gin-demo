// Package service
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/7/3 20:56
 */
package service

import (
	"fmt"
	"gin-demo/models/views"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// HelloWorld 你好世界
func (u *UserService) HelloWorld(params views.HelloVo) (vo *views.HelloResultVo, err error) {
	var result views.HelloResultVo
	patchParams(&params)
	result.Msg = fmt.Sprintf("Hello %s, beg your pardon: %s", params.Username, params.Msg)
	return &result, err
}

func patchParams(params interface{}) {
	switch params.(type) {
	case *views.HelloVo:
		helloParams := params.(*views.HelloVo)
		if helloParams.Username == "" {
			helloParams.Username = "Strange people"
		}
		if helloParams.Msg == "" {
			helloParams.Msg = "I should keep silent..."
		}

	}
}
