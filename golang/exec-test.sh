#!/bin/bash

export GOPATH=`pwd`
echo $GOPATH
#go test http-server_test.go
go test -v ./
