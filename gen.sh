#! /bin/bash

mkdir -p ./gen/thrift
# thrift
thrift -r -gen go -out ./gen/thrift/ echo.thrift

# tchannel
thrift-gen -inputFile echo.thrift -outputDir ./gen/tchannel/ -generateThrift