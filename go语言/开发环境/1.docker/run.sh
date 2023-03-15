#!/bin/bash
name=daima_docker_run
process=process
docker rm -vf $name
env CC=x86_64-unknown-linux-gnu-gcc CXX=x86_64-unknown-linux-gnu-g++ GOARCH=amd64 GOOS=linux  go build -o $process main.go
docker run -v $(pwd)/$process:/$process --name $name alpine /$process --who=chihuo
