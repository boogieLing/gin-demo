#!/bin/bash
# ===========================外部解析yaml的工具=============================
# Based on https://gist.github.com/pkuczynski/8665367
function parse_yaml() {
    local yaml_file=$1
    local prefix=$2
    local s
    local w
    local fs

    s='[[:space:]]*'
    w='[a-zA-Z0-9_.-]*'
    fs="$(echo @ | tr @ '\034')"

    (
        sed -e '/- [^\“]'"[^\']"'.*: /s|\([ ]*\)- \([[:space:]]*\)|\1-\'$'\n''  \1\2|g' |
            sed -ne '/^--/s|--||g; s|\"|\\\"|g; s/[[:space:]]*$//g;' \
                -e 's/\$/\\\$/g' \
                -e "/#.*[\"\']/!s| #.*||g; /^#/s|#.*||g;" \
                -e "s|^\($s\)\($w\)$s:$s\"\(.*\)\"$s\$|\1$fs\2$fs\3|p" \
                -e "s|^\($s\)\($w\)${s}[:-]$s\(.*\)$s\$|\1$fs\2$fs\3|p" |
            awk -F"$fs" '{
            indent = length($1)/2;
            if (length($2) == 0) { conj[indent]="+";} else {conj[indent]="";}
            vname[indent] = $2;
            for (i in vname) {if (i > indent) {delete vname[i]}}
                if (length($3) > 0) {
                    vn=""; for (i=0; i<indent; i++) {vn=(vn)(vname[i])("_")}
                    printf("%s%s%s%s=(\"%s\")\n", "'"$prefix"'",vn, $2, conj[indent-1], $3);
                }
            }' |
            sed -e 's/_=/+=/g' |
            awk 'BEGIN {
                FS="=";
                OFS="="
            }
            /(-|\.).*=/ {
                gsub("-|\\.", "_", $1)
            }
            { print }'
    ) <"$yaml_file"
}

function unset_variables() {
    # Pulls out the variable names and unsets them.
    #shellcheck disable=SC2048,SC2206 #Permit variables without quotes
    local variable_string=($*)
    unset variables
    variables=()
    for variable in "${variable_string[@]}"; do
        tmpvar=$(echo "$variable" | grep '=' | sed 's/=.*//' | sed 's/+.*//')
        variables+=("$tmpvar")
    done
    for variable in "${variables[@]}"; do
        if [ -n "$variable" ]; then
            unset "$variable"
        fi
    done
}

function create_variables() {
    local yaml_file="$1"
    local prefix="$2"
    local yaml_string
    yaml_string="$(parse_yaml "$yaml_file" "$prefix")"
    unset_variables "${yaml_string}"
    eval "${yaml_string}"
}
# ===========================外部解析yaml的工具=============================
# ===============================处理逻辑==================================
# sudo chmod +x ./docker-build.sh && sudo ./docker-build.sh -m mongod.conf -g ./config/config.yml
#if [ -f "Dockerfile.old" ];then
#  # 如果Dockerfile副本存在，用原始副本替换
#  echo "use Dockerfile.old"
#  cp -p Dockerfile.old Dockerfile
#else
#  echo "create Dockerfile.old"
#  # 后续需要修改Dockerfile，保存一个副本
#  cp -p Dockerfile Dockerfile.old
#fi
# 处理从命令行中传递的参数
mongod_config_name="mongod.conf"
mongo_db="gin_db"
mongo_db_username="super_admin"
mongo_db_password="super_password"
gin_config="./config/config.yml"
while getopts ":m:g:" opt
do
    case $opt in
        m)
        mongod_str=$OPTARG

        OLD_IFS="$IFS"
        IFS="+"
        params=($mongod_str)
        IFS="$OLD_IFS"
        for (( i=0;i<${#params[@]};i++ ))
        do
          if [ $i -eq 0 ];then
            mongod_config_name=${params[$i]}
          fi
          if [ $i -eq 1 ];then
            mongo_db_username=${params[$i]}
          fi
          if [ $i -eq 2 ];then
            mongo_db_password=${params[$i]}
          fi
        done
        ;;
        g)
        gin_config=$OPTARG
        echo "gin config: ${gin_config}"
        ;;
        ?)
        echo "unknown option: $opt"
        exit 1;;
    esac
done


files=("go1.17.11.linux-amd64.tar.gz" "mongodb-linux-x86_64-ubuntu1804-5.0.9.tgz")
files_trans=("\"go1.17.11.linux\-amd64.tar.gz\"" "\"mongodb\-linux-x86_64\-ubuntu1804\-5.0.9.tgz\"")
files_from_wget=("https:\/\/golang.google.cn\/dl\/go1.17.11.linux\-amd64.tar.gz" "https:\/\/fastdl.mongodb.org\/linux\/mongodb\-linux\-x86_64\-ubuntu1804\-5.0.9.tgz")
wget_files=()
copy_files=()
copy_flag="false"
copy_apply_atr=""
wget_flag="false"
wget_apply_str=""

# 读取文件列表，不存在的文件就加入wget队列，存在就加入copy队列
for (( i=0;i<${#files[@]};i++ ))
do
  if [ ! -f ${files[$i]} ];then
    echo -e "\033[41;30m XXX ${files[$i]} not exist ?? XXX \033[0m"
    wget_flag="true"
    wget_files[${#wget_files[*]}]=${files_from_wget[$i]}
  else
    echo -e "\033[46;30m √√√ ${files[$i]} exist !! √√√ \033[0m"
    copy_flag="true"
    copy_files[${#copy_files[*]}]=${files_trans[$i]}
  fi
done

#如果存在需要wget的文件，需要安装wget
if [ "${wget_flag}"x == "true"x ];then
  wget_apply_str="RUN apt\-get install \-\-assume-yes apt\-utils \&\& apt\-get install wget \-y"
  for wget in ${wget_files[@]}
  do
    connect_str_prefix=" \&\& wget "
    connect_str_suffix=" \-P \/tmp "
    wget_apply_str=$wget_apply_str$connect_str_prefix$wget$connect_str_suffix
  done
else
  wget_apply_str=""
fi
sed -i "s/\${LOAD_FILE_WGET}/$wget_apply_str/g" ./Dockerfile

#如果存在需要COPY的文件，提供docker的COPY指令
if [ "${copy_flag}"x == "true"x ];then
  connect_str_prefix="COPY ["
  for copy in ${copy_files[@]}
  do
    connect_str=", "
    copy_apply_atr=$copy_apply_atr$copy$connect_str
  done
  connect_str_suffix="\"\/tmp\/\"]"
  copy_apply_atr=$connect_str_prefix$copy_apply_atr$connect_str_suffix
else
  copy_apply_atr=""
fi
sed -i "s/\${LOAD_FILE_COPY}/${copy_apply_atr}/g" ./Dockerfile
# 指定mongod配置文件
sed -i "s/\${MONGOD_FILE_NAME}/${mongod_config_name}/g" ./Dockerfile
# yaml解析脚本
# 解析mongod配置的变量
create_variables ${mongod_config_name} "mongod_"
# 将mongod的端口暴露
sed -i "s/\${MONGOD_EXPOSE}/${mongod_net_port}/g" ./Dockerfile
# 解析gin配置的变量
create_variables ${gin_config} "gin_"
mongo_db=${gin_mongo_db}
if [[ ${gin_mongo_username} != "" ]]; then
    mongo_db_username=${gin_mongo_username}
fi
if [[ ${gin_mongo_password} != "" ]]; then
    mongo_db_password=${gin_mongo_password}
fi
touch tmp-mongo-add-admin.js
echo "conn = new Mongo();db = conn.getDB(\"${mongo_db}\");db.createUser({user:\"${mongo_db_username}\",pwd:\"${mongo_db_password}\",roles: [{ role: \"readWrite\", db: \"${mongo_db}\" }]});db.createCollection(\"hello\");" > tmp-mongo-add-admin.js
cat tmp-mongo-add-admin.js
# 将gin的端口暴露
# 这里截断了${gin_system_part}的第0位，因为配置文件中的端口值带了":"的前缀
sed -i "s/\${GIN_EXPOSE}/${gin_system_part:1}/g" ./Dockerfile

echo "mongod config: ${mongod_config_name}"
echo "mongod db: ${mongo_db}"
echo "mongod username: ${mongo_db_username}"
echo "mongod password: ${mongo_db_password}"
echo "docker subnet: ${gin_docker_subnet}"
echo "docker gateway: ${gin_docker_gateway}"
echo "gin config: ${gin_config}"

if [ "${gin_mongo_part}"x != "${mongod_net_port}"x ];then
  echo -e "\033[41;30m mongo port in ${gin_config}:${gin_mongo_part} not equal in ${mongod_config_name}:${mongod_net_port} \033[0m"
  exit
fi

# cat Dockerfile
# 清除旧镜像和旧网络
docker container stop gin-mongo5-mongo; docker container rm gin-mongo5-mongo ; docker image rm gin-mongo5-mongo;
docker container stop gin-mongo5-gin; docker container rm gin-mongo5-gin ; docker image rm gin-mongo5-gin;
docker network rm gin-mongo5-net;
# 构建镜像
docker build --target mongo_apply -t gin-mongo5-mongo . &&
docker build --target gin_apply -t gin-mongo5-gin . &&
# 以网桥模式新建一个docker网络
docker network create -d bridge --subnet ${gin_docker_subnet} --gateway ${gin_docker_gateway} gin-mongo5-net &&
# 以脱离模式运行容器
docker run --name gin-mongo5-mongo -d -p ${mongod_net_port}:${mongod_net_port} --network gin-mongo5-net gin-mongo5-mongo:latest &&
docker run --name gin-mongo5-gin -d -p ${gin_system_part:1}:${gin_system_part:1} --network gin-mongo5-net gin-mongo5-gin:latest
# 删除临时生成的文件
rm tmp-mongo-add-admin.js
# check
docker container ls -a
docker network ls