# 杀死进程
function kill_process() {
    for i; do
        pid=$(ps -ef | grep $i | grep -v grep | awk '{print $2}')
        if [ -n "$pid" ]; then
            kill -9 $pid
        fi
    done
}

kill_process gtools
echo "kill gtools success"

if [ ! -d "logs" ]; then
    mkdir -p logs
fi
# go build -o gtools -ldflags "-s -w"
# echo "build gtools success"

export MODE_ENV=prod
nohup ./gtools >logs/gtools.log 2>&1 &
echo "start gtools success"
