// Package utils
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description:
 * @File:  pid_record_util
 * @Version: 1.0.0
 * @Date: 2022/7/10 22:07
 */
package utils

import (
	"bufio"
	"fmt"
	"os"
)

// RecordRunPid 记录进程pid
func RecordRunPid() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("记录进程pid err=", err)
		}
	}()
	fileName := "runtime.servicePid" //记录pid文件位置（基于项目位置）
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("创建pid文件 err=", err)
	}
	pid := os.Getpid()
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(fmt.Sprintf("%v", pid))
	_ = writer.Flush()
}
