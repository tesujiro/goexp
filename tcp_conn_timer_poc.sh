#!/bin/bash

IP_ADDR=`ifconfig eth0|grep "inet addr"|awk '{gsub(/.*:/,"",$2);print $2;}'`
PORT=22
#OUTFILE=/tmp/trace_conn.$$
OUTFILE=/tmp/trace_conn.log
TIMER=5

dump() {
	sudo tcpdump -l -tt "src $IP_ADDR and src port $PORT and \
		(  tcp[tcpflags] & (tcp-syn|tcp-ack)==(tcp-syn|tcp-ack) \
		or tcp[tcpflags] & (tcp-fin|tcp-ack)==(tcp-fin|tcp-ack) )"
}

summary() {
	awk 'NF>0 {
		if ($7 ~/F/) FLAG="FIN";
		else FLAG="SYN"
		print $1,$3,$5,FLAG
		fflush()
	}'
}

list(){
	local input=$1
	local n=$2
	local t=$TIMER

	awk -v now=$n -v timer=$t '
	$4=="SYN"{
		START[$3]=$1
		if (DURATION[$3]!="") delete DURATION[$3]
	}
	$4=="FIN"{
		if (START[$3]!="") {
			DURATION[$3]=$1-START[$3]
		}
	}
	END{
		for (key in START){
			if (DURATION[key]!="") {
				if (START[key] > now - timer || START[key]+DURATION[key] > now - timer ) 
					print key,DURATION[key];
			}
			else{
				if (START[key] != "")
					print key,now - START[key];
			}
		}
	}' $input
}

#dump | summary >> $OUTFILE &
#exit
#dump | summary | display_time

while :
do
	now=`date +%s.%N`
	echo \[$now\]
	list $OUTFILE $now
	echo
	sleep $TIMER
done

