#!/bin/bash

if [ $# -ne 2 ];then
	echo PARAMETER ERROR
	exit 1
fi

PORT=$1
INTERVAL=$2
IP_ADDR=`ifconfig eth0|grep "inet "|awk '{gsub(/.*:/,"",$2);print $2;}'`

#echo IP_ADDR=$IP_ADDR
#exit

dump() {
	sudo tcpdump -l -tt "src $IP_ADDR and src port $PORT and \
		(  tcp[tcpflags] & (tcp-syn|tcp-ack)==(tcp-syn|tcp-ack) \
		or tcp[tcpflags] & (tcp-fin|tcp-ack)==(tcp-fin|tcp-ack) )"
}


dump|./connection_monitor -interval $INTERVAL

