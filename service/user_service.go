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
	"context"
	"fmt"
	"gin-demo/global"
	"gin-demo/models/views"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// HelloWorld 你好世界
func (u *UserService) HelloWorld(params views.HelloVo) (vo *views.HelloResultVo, err error) {
	var result views.HelloResultVo
	patchParams(&params)
	insertResult, err := global.DBEngine.Collection("hello").InsertOne(context.TODO(), params)
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}

	result.Msg = fmt.Sprintf("Hello %s, beg your pardon: %s, _id: %s",
		params.Username, params.Msg, insertResult.InsertedID.(primitive.ObjectID).Hex(),
	)
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
