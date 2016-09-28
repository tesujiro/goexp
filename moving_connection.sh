#!/bin/sh

INPUT_LINES=100000
#INPUT_LINES=5

log_cat()
{
	cat <<EOF | awk 'sub(/#.*/,"")>=0&&NF'
10.10.10.1 - - [19/Sep/2016:06:29:59.123 +0900] "GET /Detail HTTP/1.1" 200 15110 2.001
10.10.10.1 - - [19/Sep/2016:06:30:15.123 +0900] "GET /Detail HTTP/1.1" 200 15110 3.001
10.10.10.1 - - [19/Sep/2016:06:30:24.011 +0900] "GET /Detail HTTP/1.1" 200 15110 0.001
10.10.10.1 - - [19/Sep/2016:06:30:24.015 +0900] "GET /Detail HTTP/1.1" 200 15110 10.001
10.10.10.1 - - [19/Sep/2016:06:30:25.000 +0900] "GET /Detail HTTP/1.1" 200 15110 2.000
EOF
}

access()
{
	local loop=$(( $INPUT_LINES/$(log_cat|wc -l) ))
	local log=$(log_cat)
	for i in $(seq $loop); do
		echo "$log"
	done
}

moving_connection_awk()
{
	gawk -v T_START=$1 -v T_END=$2 -v UNIT=$3 '
	function t2u(t){
		YYYY=substr(t,8,4)
		MM=M2N[substr(t,4,3)]
		DD=substr(t,1,2)
		hh=substr(t,13,2)
		mm=substr(t,16,2)
		ss=substr(t,19,2)
		ms="0"substr(t,21,4)+0
		param=sprintf("%d %d %d %d %d %d",YYYY,MM,DD,hh,mm,ss)
		return (mktime(param)+ms)
	}
	function u2t(u){
		YMDhms=strftime("%Y/%m/%d %H:%M:%S",sprintf("%.1f",u))
		milli=substr(sprintf("%.3f",u-int(u)),3)
		return sprintf("%s.%s",YMDhms,milli)
	}
	function ceil(num) {
		if (int(num) == num) {
			return num;
		} else if (num > 0) {
			return int(num) + 1;
		} else {
			return num;
		}
	}
	BEGIN{
		CONVFMT="%.12g";
		OFMT="%.12g";

		MONTH_STRING="Jan Feb Mar Apr May Jun Jul Aug Sep Oct Nov Dec"
		split(MONTH_STRING,MONTH)
		for (i in MONTH){ M2N[MONTH[i]]=i; }

		U_START=t2u(T_START)
		U_END=t2u(T_END)
		#print "U_START=" U_START
		#print "U_END=" U_END
		#print "UNIT=" UNIT
		for(u=U_START;u<=U_END;u+=UNIT){
			count[u]=0
			#printf "u=%f\n",u
		}
	}
	{
		t_start=substr($4,2)
		u_start=t2u(t_start)
		response=$11
		#print $0
		#print "u_start=" u_start
		#print "response=" response
		#if ( ceil(u_start/UNIT)*UNIT == int((u_start+response)/UNIT)*UNIT) next;
		for(delta=0;delta<=response;delta+=UNIT){
			m=ceil((u_start+delta)/UNIT)*UNIT
			if ( m >= U_START && m <= U_END && m <= u_start + response ){
				count[m]++
				#print "count["m"]="count[m]
			}
		}
	}
	END{
		for(u=U_START;u<=U_END;u+=UNIT){
			printf "%s %d\n",u2t(u),count[u]
		}
		
	}'
}

START="19/Sep/2016:06:30:00"
END="19/Sep/2016:06:31:00"
main_go()
{
#Go
access | ./moving_connection  -start "$START +0900" -end "$END +0900" -unit 1000
}

main_awk()
{
#AWK
access | moving_connection_awk $START $END 1
}

time main_go
time main_awk

exit

if [ $# -ne 3 ];then
	echo parameter error
	exit
fi

access | moving_connection_awk $@

