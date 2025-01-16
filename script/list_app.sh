#!/usr/bin/env bash

readonly dir="app"  # 定义 app 文件夹的路径

app_list=()

get_app_list(){
    local idx=0
    for d in "$dir"/*; do
        if [ -d "$d" ] && [[ "$d" != "$dir/frontend" ]]; then
            app_list[idx]="$d"
            idx=$((idx + 1))
        fi
    done
}