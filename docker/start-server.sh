#!/bin/sh
# docker 

if [[ "$1" == "provider" ]]; then
  echo "Starting provider server..."
  cd /www/app/
  ./server.exe
else
  echo "Unrecognized arguments, by default starting provider server ."
  cd /www/app/
  ./server.exe -Name=$1
fi
