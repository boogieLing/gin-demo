package initialize

import (
	"context"
	"fmt"
	"gin-demo/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// InitDB 获取一个mongo的db操作实例
// 一定要记得defer关闭client
func InitDB() (context.Context, *mongo.Client) {
	// FOLLOW:
	// https://www.mongodb.com/docs/drivers/go/current/fundamentals/auth/#std-label-golang-authentication-mechanisms
	credential := options.Credential{
		AuthSource: global.Config.Mongo.DB,
		Username:   global.Config.Mongo.Username,
		Password:   global.Config.Mongo.Password,
	}
	dsn := fmt.Sprintf("mongodb://%s:%s/?authSource=admin",
		global.Config.Mongo.Address,
		global.Config.Mongo.Part,
	)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn).SetAuth(credential))
	if err != nil {
		global.Logger.Error(err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		global.Logger.Error(err)
	}
	// db := client.Database(global.Config.Mongo.DB)
	global.Logger.Infof("=== Connected MongoDB:%s ===", global.Config.Mongo.DB)
	return ctx, client
}
