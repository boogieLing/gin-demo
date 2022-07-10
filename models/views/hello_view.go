// Package views
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description:
 * @File:  hello_view
 * @Version: 1.0.0
 * @Date: 2022/7/10 13:06
 */
package views

type HelloVo struct {
	Username string `json:"username"`
	Msg      string `json:"msg"`
}

type HelloResultVo struct {
	Msg string `json:"msg"`
}
