// Package config
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 配置项
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/7/3 18:21
 */
package config

type SystemConfig struct {
	System `yaml:"system"`
	Logger `yaml:"logger"`
	Author `yaml:"author"`
}

type System struct {
	Part   string `yaml:"part"`   // 地址
	Status string `yaml:"status"` // 状态
}

type Logger struct {
	Path string `yaml:"path"` // 日志文件路径
}

type Author struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}
