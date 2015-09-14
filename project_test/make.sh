#!/bin/sh

filepath=$(cd "$(dirname "$0")"; pwd)
temp=$GOPATH
export GOPATH=$filepath:$GOPATH
cd $filepath/bin
go build calc
#export GOPATH=$temp

