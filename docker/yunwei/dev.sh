#!/bin/sh
echo "这是本地测试脚本"
sleep 2

echo "创建日志目录"
cd ../../
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./docker/server .
cd ./docker/yunwei


# Provider
echo "启动/删除 provider"
for(( i = 80;i <= 80; i++));
do
    echo "第一：删除provider: $i";
    docker stop  provider-$i
    docker rm -f provider-$i

    if [[ "$1" != "stop" ]]; then
        echo "第二：启动provider: $i";
        # 单个容器限制使用0.5个核心 1G内存
        docker run  -p $i:3000 --restart=always  --name provider-$i \
                    --cpus=2  -m 3g \
                    --add-host code.xueersi.com:10.97.15.72  \
                    --user codeapp:codeapp \
                    -v $(pwd)/../.env.yml:/www/app/.env.yml \
                    -v $(pwd)/../server:/www/app/server.exe \
                    -d xes-code/go-crazy $i
                    # -v $(pwd)/../include/cpp:/usr/local/include:ro \
                    # -v $(pwd)/../include/py/xes/:/usr/local/lib/python3.5/dist-packages/xes/:ro \
    fi
done
