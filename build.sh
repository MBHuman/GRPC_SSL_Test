#!/bin/bash

protoc --go_out=./proto --go-grpc_out=./proto ./proto/test.proto
