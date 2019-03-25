user  nginx;
worker_processes  4;
# worker_cpu_affinity  00000011 00001100 11000000;

error_log  off;#/var/log/nginx/error.log warn;
pid        /var/run/nginx.pid;

events {
    use epoll;
    multi_accept off;
    worker_connections  65000;
}

http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    #access_log  /var/log/nginx/access.log  main;
    access_log off;
    sendfile        on;
    #tcp_nopush     on;

    keepalive_timeout  300;

    #gzip  on;

    # include /etc/nginx/conf.d/*.conf;
}


stream {
    upstream kk5 {
	#least_conn ;
	#least_time  first_byte ; 
        zone providers 400m;
        #weigth参数表示权值，权值越高被分配到的几率越大
        #本机上的Squid开启3128端口    max_fails=5 fail_timeout=1s

        {{range .}}    
        server {{.Name}}      weight={{.Weight}}  max_fails=0 fail_timeout=1s max_conns={{.MaxConns}};
        {{end}}  
        #server provider-large:30002     weight=600 max_fails=0 fail_timeout=1s max_conns=202;
       	#server provider-medium:30001    weight=500 max_fails=0 fail_timeout=1s max_conns=195;
        #server provider-small:30000     weight=100 max_fails=0 fail_timeout=1s max_conns=120;
    }
    server {
        listen 20000;
        proxy_connect_timeout 3s;
        proxy_timeout 3s;
        proxy_pass kk5;
    }
}
