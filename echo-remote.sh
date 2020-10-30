#! /bin/bash

if [ ! -d "gen-go-13" ];then
	echo "create path: gen-go-13"
	mkdir gen-go-13
fi
thrift -r -gen go -out gen-go-13 echo.thrift
cd ./gen-go-13/echo
if [ ! -f "go.mod" ];then 
	echo "creaet go package echo:"
	go mod init echo  # or modify echo-remote.go "echo"=>"localhost.com/thrift-go/gen-go-13/echo"
fi

echo "run echo-remote client:"
# -P 指定协议
# echo 是方法名
# {1:{\"str\":\"hello\"}}  是请求的结构体:
# 	1: 代表第一个字段
#	{\"str\":\"hello\"} 是 字段类型和值
go run ./echo-remote/echo-remote.go  -p 9898 -P compact echo  {1:{\"str\":\"hello\"}}

rm go.mod
rm go.sum

# run gen-client
# go run ./gen-client/echo-remote.go  -p 9898 -P compact echo  {1:{\"str\":\"hello\"}}