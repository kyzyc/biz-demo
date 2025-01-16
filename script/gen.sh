#!/usr/bin/env bash

. script/list_app.sh

get_app_list

# 遍历 app_list 数组
for full_path in "${app_list[@]}"; do
    # 使用 basename 提取目录名
    svcName=$(basename "$full_path")

    pushd "$full_path" || exit
    go mod init github.com/kyzyc/biz-demo/"$full_path"
    cwgo client -I ../../idl --type RPC --service "$svcName" --module github.com/kyzyc/biz-demo/"$full_path" --idl ../../idl/"$svcName".proto
    cwgo server -I ../../idl --type RPC --service "$svcName" --module github.com/kyzyc/biz-demo/"$full_path" --idl ../../idl/"$svcName".proto

    popd || exit
done