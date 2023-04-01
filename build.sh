#!/bin/bash 

GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/go-hello main.go

docker build -f Dockerfile -t kong-with-plugin-image:latest .