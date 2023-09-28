#!/bin/bash

echo $1

mkdir -p "$1/cmd"

nvim -O "$1/cmd/code.go" "$1/cmd/code_test.go" "$1/cmd/inp.txt"
