FROM alpine:3.7

# 设置源地址 解决2018-04-15 左右官方资源镜像错误问题
RUN sed -i -e "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
    sed -i -e "s/http/https/" /etc/apk/repositories && apk update

# c/c++编译和运行的基本依赖   coreutils解决stdbuf
# RUN apk add --no-cache  \
#     gcc \
#     g++ \
#     libc-dev \
#     coreutils \
#     clang

# 增加独立用户--增强用户级别隔离
# ensure codeapp user exists
RUN set -x \
    && addgroup -g 82 -S codeapp \
    && adduser -u 82 -D -S -G codeapp codeapp
# 82 is the standard uid/gid for "www-data" in Alpine

ADD start-server.sh /usr/local/bin

# && chown -R codeapp:codeapp /www/app
RUN mkdir /www && mkdir /www/app 

RUN mkdir /data && mkdir /data/user && chown -R codeapp:codeapp /data/user

RUN mkdir /www/app/storage && mkdir /www/app/storage/logs && chown -R codeapp:codeapp /www/app/storage

RUN mkdir /usr/local/include

# ADD . /www/app
ADD ./.env.yml /www/app
ADD ./server.exe /www/app
ADD ./start-server.sh /www/app

RUN chmod a+x /usr/local/bin/start-server.sh && chmod a+x /www/app/server.exe

USER codeapp:codeapp

WORKDIR /tmp

VOLUME /root/logs

EXPOSE 3000

ENTRYPOINT ["start-server.sh"]