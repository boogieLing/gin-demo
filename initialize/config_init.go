// Package initialize
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description:
 * @File:  config_init
 * @Version: 1.0.0
 * @Date: 2022/7/10 12:54
 */
package initialize

import (
	"gin-demo/config"
	"gin-demo/global"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func InitDevConfig() *config.SystemConfig {
	return InitProdConfig(global.YAMLPATH)
}

// InitProdConfig 读取yml文件
func InitProdConfig(path string) *config.SystemConfig {
	r := &config.SystemConfig{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Fatalf("read config.yml error: %s\n", err.Error())
		panic(err.Error())
	}
	if yaml.Unmarshal(file, r) != nil {
		logrus.Fatalf("analysis config.yml error: %s\n", err)
		panic(err.Error())
	}
	return r
}
