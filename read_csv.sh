#!/bin/bash

echo == NORMAL CASE ==
tee <<EOF | ./read_csv
aaa,bbb,ccc
xxx,yyy,zzz
111,222,333
EOF

echo == QUOTED CASE ==
tee <<EOF | ./read_csv
"aaa","bbb","ccc"
"aa""a","bbb","ccc"
EOF
