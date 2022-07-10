#!/bin/bash
# nohup /home/ling/gin-demo/gin-server >nohup.log 2>&1 &
# 获取脚本所在目录
File=$(cd "$(dirname "$0")";pwd)

# shellcheck disable=SC2013
# shellcheck disable=SC2006
for line in `cat "$File"/runtime.servicePid`
do
  echo "Restart $line."
  kill -1 "$line"
done