#!/bin/bash
# ===========================容器启动后的操作=============================
go build -o /home/gin-server/gin-server /home/gin-server/main.go &&
nohup /home/gin-server/gin-server >nohup.log 2>&1 & &&
/usr/local/mongodb/bin/mongod -f /usr/local/mongodb/${1}