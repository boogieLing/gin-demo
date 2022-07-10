// Package middleware
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 日志打印的中间件
 * @File:  logger
 * @Version: 1.0.0
 * @Date: 2022/7/3 18:41
 */
package middleware

import (
	"fmt"
	"gin-demo/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type GINLogFormatter struct{}

// Format 控制格式
func (box *GINLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("[GIN   :%s] : [%s] - [%s]\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func Logger() gin.HandlerFunc {
	file, err := os.OpenFile(global.Config.Logger.Path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("err:", err)
	}
	logger := logrus.New()
	logger.Out = file
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(new(GINLogFormatter))
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		runTime := endTime.Sub(startTime)
		method := c.Request.Method
		url := c.Request.RequestURI
		status := c.Writer.Status()
		ip := c.ClientIP()
		logger.Infof(" %3d | %13v | %15s | %s | %s ", status, runTime, ip, method, url)
	}
}
