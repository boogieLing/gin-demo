// Package msg
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 控制Api返回的msg构造
 * @File:  ApiMsg
 * @Version: 1.0.0
 * @Date: 2022/7/3 21:51
 */
package msg

import "gin-demo/global"

var CodeMsg = map[string]string{
	global.SUCCESS: "操作成功",
	global.FAILED:  "操作失败",
}

// ApiMsg 消息实体
type ApiMsg struct {
	// code
	Code string `json:"code"`
	// msg
	Msg string `json:"msg"`
	// data
	Data interface{} `json:"data"`
}

func NewMsg() *ApiMsg {
	return new(ApiMsg)
}

// Success 成功
func (msg *ApiMsg) Success(data interface{}) *ApiMsg {
	msg.Code = global.SUCCESS
	msg.Msg = CodeMsg[global.SUCCESS]
	msg.Data = data
	return msg
}

// Failed 失败
func (msg *ApiMsg) Failed(detailMsg string) *ApiMsg {
	msg.Code = global.FAILED
	msg.Msg = CodeMsg[global.FAILED]
	msg.Data = detailMsg
	return msg
}
