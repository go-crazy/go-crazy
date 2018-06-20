
# [Wrk]
# Threads = 2
# Timeout = 5
# WarmupDuration = 30s
# PressureDuration = 60s
# SmallScale = 128
# MediumScale = 256
# LargeScale = 512

#clear
echo "Starting 测评测试..." 
git rev-parse HEAD  | tee wrk.sh.log

wrk -t2 -c512 -d30s -T5 --script=./post.lua --latency http://10.99.2.116:8087/invoke

wrk -t2 -c128 -d60s -T5 --script=./post.lua --latency http://10.99.2.116:8087/invoke   | tee wrk.sh.log

wrk -t2 -c256 -d60s -T5 --script=./post.lua --latency http://10.99.2.116:8087/invoke   | tee wrk.sh.log

wrk -t2 -c512 -d60s -T5 --script=./post.lua --latency http://10.99.2.116:8087/invoke   | tee wrk.sh.log

echo "End 测评测试..." 