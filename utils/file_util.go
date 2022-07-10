// Package utils
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 文件相关工具
 * @File:  file_utils
 * @Version: 1.0.0
 * @Date: 2022/7/5 14:15
 */
package utils

import (
	"fmt"
	"gin-demo/global"
	"os"
	"path/filepath"
	"time"
)

// todayFilename 自动获取文件名
func todayFilename() string {
	currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	path := currentPath + "/logs/" + time.Now().Format("2006-01")
	_ = os.MkdirAll(path, os.ModePerm)
	day := time.Now().Format("02")
	return path + "/" + day + ".log"
}

// newLogFile 新建当日文件对象
func newLogFile() *os.File {
	filename := todayFilename()
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error")
	}
	return file
}

// YamlLogFile yaml中配置的文件
func YamlLogFile() *os.File {
	file, err := os.OpenFile(global.Config.Logger.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("err:", err)
	}
	return file
}
