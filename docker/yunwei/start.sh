# #!/bin/sh
# echo "请注意该脚本必须在先cd到yunwei目录下执行！3"
# sleep 2
# echo "请注意该脚本必须在先cd到yunwei目录下执行！2"
# sleep 2
# echo "请注意该脚本必须在先cd到yunwei目录下执行！1"
# sleep 1
# echo ""
# echo "----------------------------------------------------"
# echo "～～欢迎使用启动脚本，如有问题请联系编程团队！～～"
# echo "----------------------------------------------------"
# echo ""

# echo "创建日志目录"
# cd ../
#     mkdir logs
# cd ./yunwei

# if [[ "$1" == "build" ]]; then
#     # echo "开始构建镜像"
#     # cd ../
#     # . build.docker.sh
#     # cd ./yunwei
#     docker rmi egistry.cn-hangzhou.aliyuncs.com/qylin-docker/codeserver:latest
#     return
# fi

# # Network
# echo "创建 network"
# # NETWORK_SUBNET = '10.11.10.0/24'
# # NETWORK_GATEWAY = '10.11.10.1'
# CID=$(docker network ls --filter name=netxescode -q)
# if [[ "$CID" == "" ]]; then
#     docker network create \
#         --driver=bridge \
#         --subnet='10.11.10.0/24' \
#         --gateway='10.11.10.1' \
#         -o "com.docker.network.bridge.name"="netxescode" \
#         -o "com.docker.network.bridge.enable_icc"="true" \
#         netxescode
# fi

# # Nginx
# echo "启动/删除 nginx"
# echo "删除 Nginx";
# docker rm -f nginx-code

# if [[ "$1" != "stop" ]]; then
#     echo "启动 Nginx";
#     mkdir ../logs/nginx
#     chown -R 82:82 ../logs/
#     docker run -p 3000:80  --restart=always  --name nginx-code --network=netxescode \
#                 -v $(pwd)/../logs/nginx:/var/log/nginx \
#                 -v $(pwd)/../nginx/nginx.conf:/etc/nginx/nginx.conf:rw \
#                 -v $(pwd)/../nginx/default.conf:/etc/nginx/conf.d/default.conf:rw \
#                 -d nginx:1.15-alpine
# fi

# # Provider
# echo "启动/删除 provider"
# for(( i = 8000;i <= 8005; i++));
# do 
#     echo "第一：删除provider: $i";
#     docker rm -f provider-$i
#     mkdir ../logs/providers
#     chown -R 82:82 ../logs/
#     chown -R 82:82 ../data/

#     if [[ "$1" != "stop" ]]; then
#         echo "第二：启动provider: $i";
#         # 单个容器限制使用0.5个核心 1G内存
#         docker run  -p $i:3000 --restart=always  --name provider-$i \
#                     --network=netxescode \
#                     --cpus=2  -m 3g \
#                     --user codeapp:codeapp \
#                     -v $(pwd)/../include/cpp:/usr/local/include:ro \
#                     -v $(pwd)/../include/py/xes/:/usr/local/lib/python3.5/dist-packages/xes/:ro \
#                     -v $(pwd)/../data:/data/user \
#                     -v $(pwd)/../logs/providers:/www/app/storage/logs \
#                     -v $(pwd)/../.env.yml:/www/app/.env.yml \
#                     -v $(pwd)/../server.exe:/www/app/server.exe \
#                     -d registry.cn-hangzhou.aliyuncs.com/qylin-docker/codeserver:online $i
#                     #--add-host code.xueersi.com:10.97.15.72  \
#     fi
# done

# #  删除网桥
# if [[ "$1" == "stop" ]]; then
#     docker network rm netxescode
# fi
