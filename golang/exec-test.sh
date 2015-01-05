#!/bin/bash

export GOPATH=`pwd`
echo $GOPATH
go test -v ./
