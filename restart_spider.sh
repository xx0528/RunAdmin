#!/bin/sh

# 停止已有的 spider 服务
pkill -f spiderService

# 等待已有服务退出
sleep 3

# 启动新的 spider 服务
nohup python3.10 /root/server/spiderService &

#编辑 crontab，设定每三个小时执行 restart_spider.sh 脚本。
#crontab -e
#在打开的编辑器中添加如下内容：
#0 */3 * * * /root/server/restart_spider.sh

#每三分钟
#*/3 * * * * /root/server/restart_spider.sh


#每天八点重启
#0 8 * * * /root/server/restart_spider.sh