// Package global
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description:
 * @File:  global
 * @Version: 1.0.0
 * @Date: 2022/7/10 12:54
 */
package global

import (
	"gin-demo/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Config       *config.SystemConfig
	Logger       *logrus.Logger
	ClientEngine *mongo.Client
	DBEngine     *mongo.Database
)
