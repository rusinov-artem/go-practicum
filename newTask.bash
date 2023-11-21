#!/bin/bash

codeFile="$1/cmd/code.go"
testFile="$1/cmd/code_test.go"
inpFile="$1/cmd/inp.txt"

codeFileContent=$(
cat << EOF
package main

EOF
)

testFileContent=$(
cat << EOF
package main

 import (
     "testing"

    "github.com/stretchr/testify/assert"
 )

 func Test_(t *testing.T) {
    assert.True(t, true) 
 }
EOF
)

mkdir -p "$1/cmd"
if ! [ -f ${codeFile} ]; then
    touch "$codeFile"
    echo  "${codeFileContent}" > ${codeFile}
fi

if ! [ -f ${testFile} ]; then
    touch "$testFile"
    echo "${testFileContent}" > ${testFile}
fi

touch ${inpFile}

bdir=`basename $1 | sed 's/\./_/g'`

tmux new-session -c $1 -d -s $bdir \
    && tmux switch-client -t $bdir || tmux new -c $1 -A -s $bdir && \
   sleep 1 && \
   tmux send-keys -t ${bdir}:1.1  "nvim -O cmd/code.go cmd/code_test.go cmd/inp.txt" Enter

