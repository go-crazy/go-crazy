#!/bin/bash

# Old Version
# ETCD_HOST=$(ip addr show docker0 | grep 'inet\b' | awk '{print $2}' | cut -d '/' -f 1)
# ETCD_PORT=2379

# if [ "$ETCD_HOST"x == ""x ]
# then
#    ETCD_URL=http://172.17.0.1:$ETCD_PORT
# else
#    ETCD_URL=http://$ETCD_HOST:$ETCD_PORT
# fi

# docker 
ETCD_HOST=etcd
ETCD_PORT=2379
ETCD_URL=http://$ETCD_HOST:$ETCD_PORT

# ETCD_URL=http://172.17.0.1:$ETCD_PORT

echo ETCD_URL = $ETCD_URL

if [[ "$1" == "consumer" ]]; then
  echo "Starting consumer agent..."
  cd /root/workspace/agent/
    ./server.exe -DConnSize=1  -Dtype=consumer -Dserver.port=20000  -Detcd.url=$ETCD_URL  -Dlogs.dir=/root/logs 

else
  echo "Unrecognized arguments, exit."
  exit 1
fi
