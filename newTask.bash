#!/bin/bash

echo $1

mkdir -p "$1/cmd"

bdir=`basename $1 | sed 's/\./_/g'`

tmux new-session -c $1 -d -s $bdir \
    && tmux switch-client -t $bdir || tmux new -c $1 -A -s $bdir #&& \
   tmux send-keys -t ${bdir}:1.1  "nvim -O \"cmd/code.go\" \"cmd/code_test.go\" \"cmd/inp.txt\"" Enter  

