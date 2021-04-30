#!/bin/bash

rm -f ../src/gonetlockprotobuf/protobuf/*
protoc --proto_path=./ --go_out=../src/gonetlockprotobuf/protobuf --go_opt=paths=source_relative *.proto
