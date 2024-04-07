#!/bin/bash

# 定义程序名称和源代码路径
PROGRAM_NAME="go-ocr"

# 停止正在运行的 go-ocr 程序
if pgrep -f "$PROGRAM_NAME" > /dev/null; then
    echo "Stopping $PROGRAM_NAME..."
    pkill -f "$PROGRAM_NAME"
fi

# 检查是否停止成功
sleep 2 # 等待一小段时间让程序正常停止
if pgrep -f "$PROGRAM_NAME" > /dev/null; then
    echo "Error: Failed to stop $PROGRAM_NAME."
else
    echo "$PROGRAM_NAME has been stopped."
fi

go build -ldflags "-w -s" -o $PROGRAM_NAME main.go

# 检查编译是否成功
if [ $? -eq 0 ]; then
    echo "Compilation of $PROGRAM_NAME was successful."

    # 使用 nohup 启动新编译的程序
    echo "Starting $PROGRAM_NAME in background with nohup..."
    nohup ./$PROGRAM_NAME 2>&1 &
    echo "The process ID of the new $PROGRAM_NAME instance is $!"
else
    echo "Error: Compilation of $PROGRAM_NAME failed."
fi
