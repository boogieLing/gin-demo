FROM ubuntu:18.04 AS os_file_apply
LABEL maintainer="r0"
# 升级 apt-get
RUN sed -i "s@/archive.ubuntu.com/@/mirrors.tuna.tsinghua.edu.cn/@g" /etc/apt/sources.list \
    && sed -i "s@/security.ubuntu.com/@/mirrors.ustc.edu.cn/@g" /etc/apt/sources.list\
    && rm -Rf /var/lib/apt/lists/* \
    && apt-get update --fix-missing -o Acquire::http::No-Cache=True \
    && apt-get install libcurl4 openssl -y

# 以wget的方式准备文件到/tmp
${LOAD_FILE_WGET}

# 以COPY的方式准备文件到/tmp
${LOAD_FILE_COPY}

# 复制当前项目所有文件
COPY ["./", "/home/gin-server"]
RUN true

FROM os_file_apply AS gin_apply
COPY --from=os_file_apply /tmp/go1.17.11.linux-amd64.tar.gz /tmp/
RUN true
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /home/gin-server
# 安装 golang1.17.11，编译
RUN set -x; tar -zxvf /tmp/go1.17.11.linux-amd64.tar.gz -C /usr/local/\
    && echo "export GOROOT=/usr/local/go" >> ~/.bashrc\
    && echo "export GOBIN=\$GOROOT/bin" >> ~/.bashrc\
    && echo "export PATH=\$GOROOT/bin:\$PATH" >> ~/.bashrc\
    && /bin/bash -c "source ~/.bashrc" \
    && rm -rf /tmp/go1.17.11.linux-amd64.tar.gz \
    && chmod +x /home/gin-server/restart-server.sh \
    && /usr/local/go/bin/go build -o ./gin-server main.go

# 暴露gin端口
# EXPOSE 8080
EXPOSE ${GIN_EXPOSE}

ENTRYPOINT ["./gin-server"]

FROM os_file_apply AS mongo_apply
# 此COPY是必要的，所以可以将额外的COPY也放到这里
COPY ["${MONGOD_FILE_NAME}", "entrypoint.sh", "tmp-mongo-add-admin.js", "/tmp/"]
RUN true
COPY --from=os_file_apply /tmp/mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz /tmp/
RUN true
# 安装 mongo
RUN tar -zxvf /tmp/mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz -C /usr/local/ \
    && mv /usr/local/mongodb-linux-x86_64-ubuntu1804-5.0.9 /usr/local/mongodb \
    && mv /tmp/${MONGOD_FILE_NAME} /usr/local/mongodb/mongod.conf\
    && echo "export PATH=\$PATH:/usr/local/mongodb/bin" >> ~/.bashrc \
    && /bin/bash -c "source ~/.bashrc"\
    && mkdir -p /usr/local/mongodb/data/db \
    && mkdir -p /usr/local/mongodb/logs \
    && touch /usr/local/mongodb/logs/mongod.log \
    && rm -rf /tmp/mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz\
    # 先以无验证的方式后台启动mongod，实际ENTRYPOINT时必须开启验证
    && /usr/local/mongodb/bin/mongod -f /usr/local/mongodb/${MONGOD_FILE_NAME} --fork --logpath /usr/local/mongodb/logs/mongod.log\
    && sleep 5\
    && /usr/local/mongodb/bin/mongo /tmp/tmp-mongo-add-admin.js

# 暴露mongod端口
# EXPOSE 27017
EXPOSE ${MONGOD_EXPOSE}

ENTRYPOINT ["/usr/local/mongodb/bin/mongod", "--auth", "-f", "/usr/local/mongodb/${MONGOD_FILE_NAME}"]

# 流程：构建->脱离模式运行容器->即时设置权限
# 构建镜像
# docker build -t gin-mongo5 .
# 以脱离模式运行容器
# docker run --name gin-mongo5 -d -p 27017:27017 gin-mongo5:latest
# 以交互终端运行容器 并覆盖entrypoint
# docker run --name gin-mongo5-gin -it -p 8202:8202 --entrypoint="/bin/bash -c \"cd /home/gin-server && go build -o /home/gin-server/gin-server /home/gin-server/main.go\"" gin-mongo5-gin:latest
# docker run --name gin-mongo5-mongo -it -p 27017:27017 --entrypoint="/bin/bash" --network gin-mongo5-net gin-mongo5-mongo:latest

# docker attach gin-mongo5
# docker exec -it gin-mongo5-mongo /bin/bash

# 停止、删除容器和镜像
# docker container stop gin-mongo5-mongo; docker container rm gin-mongo5-mongo ; docker image rm gin-mongo5-mongo
# docker container stop gin-mongo5-gin; docker container rm gin-mongo5-gin ; docker image rm gin-mongo5-gin
# docker network rm gin-mongo5-net
