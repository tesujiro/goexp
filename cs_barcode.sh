#!/bin/sh

list()
{
	cat <<EOF
<<<<<<< HEAD
abc
+
*
&
=======
>>>>>>> b60802bdb966fc7331d4bc15f6c947c5581f2d14
1
2
3
A
B
H
I
J
1-2-3ABA
1234567890-ABCDEFGHIJKLMNOPQR
0000000000
0000000001
0000000002
0000000003
0000000009
000000000-
000000000A
000000000B
000000000C
000000000I
000000000J
EOF
}

list | awk '
BEGIN{
	# CONTROL CODE LIST
	CC1="a";CC2="b";CC3="c";CC4="d";
	CC5="e";CC6="f";CC7="g";CC8="h";
	# START CODE/STOP CODE
	ST="(";SP=")";

<<<<<<< HEAD
	INPUT_CHAR_REGEXP="[0-9A-Z-]+"

=======
>>>>>>> b60802bdb966fc7331d4bc15f6c947c5581f2d14
	# SPECIFIED MAX LENGTH 
	MAX_LENGTH=10

	# ALPHABET CONVERSION CODE LIST
	A_CODE["A"]=CC1"0"; A_CODE["B"]=CC1"1"; A_CODE["C"]=CC1"2"; A_CODE["D"]=CC1"3"; A_CODE["E"]=CC1"4";
	A_CODE["F"]=CC1"5"; A_CODE["G"]=CC1"6"; A_CODE["H"]=CC1"7"; A_CODE["I"]=CC1"8"; A_CODE["J"]=CC1"9";
	A_CODE["K"]=CC2"0"; A_CODE["L"]=CC2"1"; A_CODE["M"]=CC2"2"; A_CODE["N"]=CC2"3"; A_CODE["O"]=CC2"4";
	A_CODE["P"]=CC2"5"; A_CODE["Q"]=CC2"6"; A_CODE["R"]=CC2"7"; A_CODE["S"]=CC2"8"; A_CODE["T"]=CC2"9";
	A_CODE["U"]=CC3"0"; A_CODE["V"]=CC3"1"; A_CODE["W"]=CC3"2"; A_CODE["X"]=CC3"3"; A_CODE["Y"]=CC3"4";
	A_CODE["Z"]=CC3"5";

	# CHARACTER CODE
	C_CODE["0"]=0; C_CODE["1"]=1; C_CODE["2"]=2; C_CODE["3"]=3; C_CODE["4"]=4;
	C_CODE["5"]=5; C_CODE["6"]=6; C_CODE["7"]=7; C_CODE["8"]=8; C_CODE["9"]=9;
	C_CODE["-"]=10;
	C_CODE[CC1]=11; C_CODE[CC2]=12; C_CODE[CC3]=13; C_CODE[CC4]=14;
	C_CODE[CC5]=15; C_CODE[CC6]=16; C_CODE[CC7]=17; C_CODE[CC8]=18;

	# CHECK DIGIT CODE ( REVERSE HASH TO C_CODE )
	for (key in C_CODE) { CD_CODE[C_CODE[key]]=key; }
}
function justify(str){
	# IF LONGER THAN or EQUAL TO MAX_LENGTH, THEN CUT TAIL
	if (length(str)>=MAX_LENGTH){
		return substr(str,1,MAX_LENGTH)
	}
	# IF LESS THAN MAX_LENGTH, THEN PAD CC4
	pad=""
	for(i=1;i<=MAX_LENGTH - length(str);i++){
		pad=pad CC4
	}
	return str pad
}
function convert_alphabet(str){
	new_str=""
	for (i=1;i<=length(str);i++){
		ch=substr(str,i,1)
		if(ch~/[A-Z]/){
			new_str=new_str A_CODE[ch]
		}
		else {
			new_str=new_str ch
		}	
	}
	return new_str
}
function parity(str){
	sum=0
	for (i=1;i<=length(str);i++){
		sum+=C_CODE[substr(str,i,1)]
	}
	digit_num=19 - sum % 19
	return CD_CODE[digit_num]
}
{
<<<<<<< HEAD
	if ($0 !~ INPUT_CHAR_REGEXP) {
		print $0,"=> ERR"
		next
	}

=======
>>>>>>> b60802bdb966fc7331d4bc15f6c947c5581f2d14
	c=convert_alphabet($0)
	j=justify(c)
	p=parity(j)
	print $0,"=>",c,"=>",j,"=>",p,"==>",ST j p SP
}'
