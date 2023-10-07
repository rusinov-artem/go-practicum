#/bin/bash

DIR=$1

cd "$DIR/cmd" && go build code.go && ./code < inp.txt && rm code 

