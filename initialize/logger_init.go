// Package initialize
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description:
 * @File:  logger_init
 * @Version: 1.0.0
 * @Date: 2022/7/10 12:58
 */
package initialize

import (
	"fmt"
	"gin-demo/global"
	"gin-demo/utils"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func InitLogger() {
	ioWriter := utils.YamlLogFile()
	global.Logger = logrus.New()
	global.Logger.SetOutput(ioWriter)
	global.Logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	global.Logger.SetFormatter(new(GlobalLogFormatter))
}

type GlobalLogFormatter struct{}

func (box *GlobalLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("<GLOBAL:%s> : [%s] - [%s]\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}
