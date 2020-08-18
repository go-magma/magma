#!/bin/sh

# set -x
# set -v
CWD=$(pwd)
cd $1
PDIR=$(pwd)
shift
protoc -I. -I/usr/include $* --go_out=plugins=grpc,Mgoogle/protobuf/field_mask.proto=google.golang.org/genproto/protobuf/field_mask,Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,paths=source_relative:$CWD ./*.proto