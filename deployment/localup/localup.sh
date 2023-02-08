#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env

bin_name=greenfield-relayer
bin=${workspace}/../../build/${bin_name}

function start() {
    size=$1
    rm -rf ${workspace}/.local
    mkdir -p ${workspace}/.local
    for ((i=0;i<${size};i++));do
        mkdir -p ${workspace}/.local/relayer${i}/logs
        nohup ${bin} run --config-type local \
           --config-path ${workspace}/../../config/local/config_local_${i}.json \
           --log_dir json > ${workspace}/.local/relayer${i}/logs/relayer.log &
    done
}

function stop() {
    ps -ef | grep ${bin_name} | awk '{print $2}' | xargs kill
}

CMD=$1
SIZE=3
if [ ! -z $2 ] && [ "$2" -gt "0" ]; then
    SIZE=$2
fi
case ${CMD} in
start)
    echo "===== start ===="
    start $SIZE
    echo "===== end ===="
    ;;
stop)
    echo "===== stop ===="
    stop
    echo "===== end ===="
    ;;
*)
esac
