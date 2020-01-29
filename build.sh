#!/usr/bin/env bash

curDir=$(
    cd $(dirname $0)
    pwd
)
cd ${curDir}


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $1 main.go

scp $1 root@192.168.0.195:/root