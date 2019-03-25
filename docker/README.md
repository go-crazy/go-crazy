# 关于打包

- app 编译

  采用本地编译,编译脚本:build.app.sh
  好处:最小化镜像, 无需在镜像中部署 go 环境

- docker build

  见 build.docker.sh

- 本地定位调试

  见 run.sh,内有完整的镜像启动脚本

- 关于测试

  参见上级 tests 目录中的 wrk 脚本。

* 运维

```
  构建镜像
  . start.sh build
  启动所有服务
  . start.sh
  停止所有服务
  . start.sh stop
```

# 特殊版本

https://www.tensorflow.org/install/pip

tensorflow 在测试环境因为 cpu 指令过低不兼容，只能使用 1.5 以下版本

--add-host 处理测试环境 host 绑定问题
