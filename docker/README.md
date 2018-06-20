# 关于打包

- app 编译

  采用本地编译,编译脚本:build.app.sh
  好处:最小化镜像, 无需在镜像中部署 go 环境

- docker build

  见 build.docker.sh

- 本地定位调试

  见 run.sh,内有完整的镜像启动脚本

-  关于测试

  参见上级 tests 目录中的 wrk 脚本。
