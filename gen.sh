#! /bin/bash

mkdir -p ./gen/thrift
# thrift
thrift_13 -r -gen go -out ./gen/thrift/ echo.thrift

# tchannel
thrift-gen -inputFile echo.thrift -outputDir ./gen/tchannel/ -generateThrift -thriftBinary thrift_9