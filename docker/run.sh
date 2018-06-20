git pull

chmod a+x docker/start-agent.sh

docker stop etcd
docker rm etcd

docker stop provider-small
docker rm provider-small
docker stop provider-medium
docker rm provider-medium
docker stop provider-large
docker rm provider-large
docker stop consumer
docker rm consumer


# docker network ls --filter name=benchmarker -q

docker run -d -p 2379:2379 --name=etcd --cpu-period=50000 --cpu-quota=20000 --memory=1g --network=benchmarker   registry.cn-hangzhou.aliyuncs.com/aliware2018/alpine-etcd
docker run -d -p 30002:30002 --cpu-period=50000 --network=benchmarker --cpu-quota=90000 -m 6g  --name provider-large -v /www/middlewar/mydemo/go-crazy/docker/start-agent.sh:/usr/local/bin/start-agent.sh   -v /www/middlewar/mydemo/go-crazy/docker/server.exe:/root/workspace/agent/server.exe middlewar/agent:1.0  provider-large
docker run -d -p 30001:30001 --cpu-period=50000 --network=benchmarker --cpu-quota=60000 -m 4g  --name provider-medium -v /www/middlewar/mydemo/go-crazy/docker/start-agent.sh:/usr/local/bin/start-agent.sh   -v /www/middlewar/mydemo/go-crazy/docker/server.exe:/root/workspace/agent/server.exe middlewar/agent:1.0  provider-medium
docker run -d -p 30000:30000 --cpu-period=50000 --network=benchmarker --cpu-quota=30000 -m 2g  --name provider-small  -v /www/middlewar/mydemo/go-crazy/docker/start-agent.sh:/usr/local/bin/start-agent.sh   -v /www/middlewar/mydemo/go-crazy/docker/server.exe:/root/workspace/agent/server.exe middlewar/agent:1.0  provider-small
docker run -p 8087:8087 -p 8088:8088 -p 20000:20000 --cpu-period=50000 --network=benchmarker --cpu-quota=180000 -m 3g   --name consumer -v /www/middlewar/mydemo/go-crazy/docker/start-agent.sh:/usr/local/bin/start-agent.sh   -v /www/middlewar/mydemo/go-crazy/docker/server.exe:/root/workspace/agent/server.exe middlewar/agent:1.0  consumer


# docker run -p 30002:30002 --cpu-period=50000 --network=benchmarker --cpu-quota=90000 -m 6g  --name provider-large -v /www/middlewar/mydemo/go-crazy/docker/start-agent.sh:/usr/local/bin/start-agent.sh   -v /www/middlewar/mydemo/go-crazy/docker/server.exe:/root/workspace/agent/server.exe middlewar/agent:1.0  provider-large
